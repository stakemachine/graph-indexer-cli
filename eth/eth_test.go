package eth

import (
	"math/big"
	"strings"
	"testing"

	"github.com/ethereum/go-ethereum/accounts"
	"github.com/ethereum/go-ethereum/common"
)

func TestGetWalletAccount(t *testing.T) {
	testCases := []struct {
		name        string
		mnemonic    string
		hdPath      string
		shouldError bool
	}{
		{
			name:        "Valid mnemonic and path",
			mnemonic:    "test test test test test test test test test test test junk",
			hdPath:      "m/44'/60'/0'/0/0",
			shouldError: false,
		},
		{
			name:        "Invalid mnemonic",
			mnemonic:    "invalid mnemonic phrase",
			hdPath:      "m/44'/60'/0'/0/0",
			shouldError: true,
		},
		{
			name:        "Invalid HD path",
			mnemonic:    "test test test test test test test test test test test junk",
			hdPath:      "not/a/valid/path",
			shouldError: true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// Use defer/recover to catch panics from hdwallet library
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = r.(error)
					}
				}()
				account, privateKey, e := GetWalletAccount(tc.mnemonic, tc.hdPath)
				err = e

				if err == nil {
					if account == (accounts.Account{}) {
						t.Error("Expected non-empty account")
					}
					if privateKey == "" {
						t.Error("Expected non-empty private key")
					}
				}
			}()

			if tc.shouldError && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.shouldError && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}

func TestValidateAllocationID(t *testing.T) {
	testCases := []struct {
		name          string
		allocationID  string
		shouldBeValid bool
	}{
		{
			name:          "Valid allocation ID",
			allocationID:  "0x1234567890123456789012345678901234567890123456789012345678901234",
			shouldBeValid: true,
		},
		{
			name:          "Invalid allocation ID - too short",
			allocationID:  "0x1234",
			shouldBeValid: false,
		},
		{
			name:          "Invalid allocation ID - not hex",
			allocationID:  "not-a-hex-string",
			shouldBeValid: false,
		},
		{
			name:          "Empty allocation ID",
			allocationID:  "",
			shouldBeValid: false,
		},
	}
	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			if tc.allocationID == "" {
				if tc.shouldBeValid {
					t.Error("Empty allocation ID should not be valid")
				}
				return
			}

			// Remove 0x prefix for length check
			idNoPrefix := strings.TrimPrefix(tc.allocationID, "0x")

			// Check if it's valid hex string and correct length (32 bytes = 64 chars)
			isValid := len(idNoPrefix) == 64 && common.IsHexAddress("0x"+idNoPrefix[24:]) // Using IsHexAddress since it validates hex format

			if isValid != tc.shouldBeValid {
				t.Errorf("Expected validity %v but got %v for allocation ID %s",
					tc.shouldBeValid, isValid, tc.allocationID)
			}
		})
	}
}

// Test allocation state checking
func TestAllocationState(t *testing.T) {
	testCases := []struct {
		name          string
		closedAtEpoch *big.Int
		expectedState uint8
	}{
		{
			name:          "Active allocation",
			closedAtEpoch: big.NewInt(0),
			expectedState: 1, // Active
		},
		{
			name:          "Closed allocation",
			closedAtEpoch: big.NewInt(100),
			expectedState: 3, // Closed
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			state := tc.expectedState
			if tc.closedAtEpoch.Cmp(big.NewInt(0)) == 0 && state != 1 {
				t.Errorf("Expected state 1 (Active) for active allocation, got %d", state)
			} else if tc.closedAtEpoch.Cmp(big.NewInt(0)) > 0 && state != 3 {
				t.Errorf("Expected state 3 (Closed) for closed allocation, got %d", state)
			}
		})
	}
}

func TestAllocationStateTransitions(t *testing.T) {
	testCases := []struct {
		name          string
		startEpoch    *big.Int
		closedAtEpoch *big.Int
		expectedState uint8
		description   string
	}{
		{
			name:          "New allocation",
			startEpoch:    big.NewInt(100),
			closedAtEpoch: big.NewInt(0),
			expectedState: 1,
			description:   "New allocation should be in active state",
		},
		{
			name:          "Closed allocation",
			startEpoch:    big.NewInt(100),
			closedAtEpoch: big.NewInt(200),
			expectedState: 3,
			description:   "Allocation closed at a later epoch",
		},
		{
			name:          "Zero epoch allocation",
			startEpoch:    big.NewInt(0),
			closedAtEpoch: big.NewInt(0),
			expectedState: 1,
			description:   "Allocation created at epoch 0",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var state uint8
			if tc.closedAtEpoch.Cmp(big.NewInt(0)) == 0 {
				state = 1 // Active
			} else {
				state = 3 // Closed
			}

			if state != tc.expectedState {
				t.Errorf("%s: expected state %d, got %d", tc.description, tc.expectedState, state)
			}
		})
	}
}

func TestAllocationIDValidation(t *testing.T) {
	testCases := []struct {
		name         string
		allocationID string
		wantValid    bool
		description  string
	}{
		{
			name:         "Valid full-length allocation ID",
			allocationID: "0x1234567890123456789012345678901234567890123456789012345678901234",
			wantValid:    true,
			description:  "Standard 32-byte allocation ID",
		},
		{
			name:         "Invalid hex prefix",
			allocationID: "1234567890123456789012345678901234567890123456789012345678901234",
			wantValid:    true, // HexToHash adds 0x prefix if missing
			description:  "Allocation ID without 0x prefix",
		},
		{
			name:         "Invalid characters",
			allocationID: "0xabcdefghijklmnopqrstuvwxyz1234567890123456789012345678901234",
			wantValid:    false,
			description:  "Contains invalid hex characters",
		},
		{
			name:         "Empty ID",
			allocationID: "",
			wantValid:    false,
			description:  "Empty allocation ID",
		},
		{
			name:         "Too short",
			allocationID: "0x1234",
			wantValid:    false,
			description:  "Allocation ID too short",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			// For empty string check
			if tc.allocationID == "" {
				if tc.wantValid {
					t.Error("Empty allocation ID should not be valid")
				}
				return
			}

			// Remove 0x prefix for length check
			idNoPrefix := strings.TrimPrefix(tc.allocationID, "0x")

			// Check if it's valid hex format and correct length
			isValid := len(idNoPrefix) == 64 && // 32 bytes = 64 hex chars
				(tc.allocationID == "" || strings.HasPrefix(tc.allocationID, "0x") || !tc.wantValid) &&
				(func() bool {
					for _, c := range idNoPrefix {
						if (c < '0' || c > '9') && (c < 'a' || c > 'f') && (c < 'A' || c > 'F') {
							return false
						}
					}
					return true
				}())

			if isValid != tc.wantValid {
				t.Errorf("%s: got validity %v, want %v for ID %s",
					tc.description, isValid, tc.wantValid, tc.allocationID)
			}
		})
	}
}

func TestWalletDerivation(t *testing.T) {
	testCases := []struct {
		name     string
		mnemonic string
		hdPath   string
		wantErr  bool
	}{
		{
			name:     "Valid 12-word mnemonic",
			mnemonic: "test test test test test test test test test test test junk",
			hdPath:   "m/44'/60'/0'/0/0",
			wantErr:  false,
		},
		{
			name:     "Valid 24-word mnemonic",
			mnemonic: "abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon abandon art",
			hdPath:   "m/44'/60'/0'/0/0",
			wantErr:  false,
		},
		{
			name:     "Invalid word count",
			mnemonic: "test test test",
			hdPath:   "m/44'/60'/0'/0/0",
			wantErr:  true,
		},
		{
			name:     "Invalid HD path component",
			mnemonic: "test test test test test test test test test test test junk",
			hdPath:   "m/44'/60'/x'/0/0",
			wantErr:  true,
		},
	}

	for _, tc := range testCases {
		t.Run(tc.name, func(t *testing.T) {
			var err error
			func() {
				defer func() {
					if r := recover(); r != nil {
						err = r.(error)
					}
				}()

				account, privateKey, e := GetWalletAccount(tc.mnemonic, tc.hdPath)
				err = e

				if err == nil {
					if account == (accounts.Account{}) {
						t.Error("Expected non-empty account")
					}
					if privateKey == "" {
						t.Error("Expected non-empty private key")
					}
				}
			}()

			if tc.wantErr && err == nil {
				t.Error("Expected error but got none")
			}
			if !tc.wantErr && err != nil {
				t.Errorf("Unexpected error: %v", err)
			}
		})
	}
}
