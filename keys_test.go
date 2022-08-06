package main

import (
	"strings"
	"testing"
)

func TestNoPrivateKey(t *testing.T) {
	_, err := LoadPrivateKey()

	got := err.Error()

	want := "failed to load private key"

	if !strings.Contains(got, want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestBadPrivateKey(t *testing.T) {
	t.Setenv("PRIVATE_KEY", "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAlwAAAAdzc2gtcn\nNhAAAAAwEAAQAAAIEAqa/P7IXQsTAzCZ3GqWDc8pfriDGDeg6Sx1OO5GdsPN/5whRM4U44\nsRcqldVoNyS/sKrsLaJbdWjG1aJb5nuFSHzMY8IHms1I34rTsIE5zAPv/DIiYJjJE22r3x\nW55OOLpRHl5/pWsGClSVKm/1RdtMK4Vbr3S3Ea1acs4QcoeXUAAAIABipskgYqbJIAAAAH\nc3NoLXJzYQAAAIEAqa/P7IXQsTAzCZ3GqWDc8pfriDGDeg6Sx1OO5GdsPN/5whRM4U44sR\ncqldVoNyS/sKrsLaJbdWjG1aJb5nuFSHzMY8IHms1I34rTsIE5zAPv/DIiYJjJE22r3xW5\n5OOLpRHl5/pWsGClSVKm/1RdtMK4Vbr3S3Ea1acs4QcoeXUAAAADAQABAAAAgCUe+gyTJZ\nltouyqoGgzcYJ8q6EKu+l6wK9sXykmWu5iTSAhEsSDOTG8kKtgmUDfADRk3/AHwZxIxJ69\nuv8L7JcYzJcLoBbo0naICkHNgroHNUqD5APG9+OvSIrM1LMx610cmRle3IxaVdemgRNGLS\n3I0h5MCVBYYXCJnh8dV+ytAAAAQQC94EFjIU/77YA5XtFStRpMaViOunPQaW7MqEgP+fHt\nQL5ZQ86w+Bvn41fJMNtFyaIq4SMhNH3pScBfv7QgubK1AAAAQQDR6t2vthsBw1q2+utFug\nsFVJvEjFZjZXpbCenmm8Vg8INGJnZcygeY86qPhpBvUhmuRX4pKbCQcsU7Hw6/aEbvAAAA\nQQDO8AaF5az9Z38tE5KddwNCSv6OgL5kiNn+ZdpydhAbemthdzww+By9AzQ6TUX26TdNFt\nisgTFay3AqFSEWgOXbAAAABm5vbmFtZQECAwQ=\n-----END OPENSSH PRIVATE KEY-----")
	_, err := LoadPrivateKey()

	got := err.Error()

	want := "x509: failed to parse EC private key"

	if !strings.Contains(got, want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestPrivateKey(t *testing.T) {
	t.Setenv("PRIVATE_KEY", "-----BEGIN EC PRIVATE KEY-----\nMHcCAQEEINNrq/yRU3zH6Q+x9DGbg2NJMhTPMspaVDrvkS+buXBLoAoGCCqGSM49\nAwEHoUQDQgAEBBVqOkBDYFlsgv8yRp+w79TMhm/IiCjwbmPk4g75wKdXpCrKvN/m\nyjfjFip4j8lVuOoi9nVrghbCejVguO3PCA==\n-----END EC PRIVATE KEY-----")

	pk, _ := LoadPrivateKey()

	if pk == nil {
		t.Error("private key not loaded")
	}
}

func TestNoPublicKey(t *testing.T) {
	_, err := LoadPublicKey()

	got := err.Error()

	want := "failed to load public key"

	if !strings.Contains(got, want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestBadPublicKey(t *testing.T) {
	t.Setenv("PUBLIC_KEY", "-----BEGIN OPENSSH PRIVATE KEY-----\nb3BlbnNzaC1rZXktdjEAAAAABG5vbmUAAAAEbm9uZQAAAAAAAAABAAAAlwAAAAdzc2gtcn\nNhAAAAAwEAAQAAAIEAqa/P7IXQsTAzCZ3GqWDc8pfriDGDeg6Sx1OO5GdsPN/5whRM4U44\nsRcqldVoNyS/sKrsLaJbdWjG1aJb5nuFSHzMY8IHms1I34rTsIE5zAPv/DIiYJjJE22r3x\nW55OOLpRHl5/pWsGClSVKm/1RdtMK4Vbr3S3Ea1acs4QcoeXUAAAIABipskgYqbJIAAAAH\nc3NoLXJzYQAAAIEAqa/P7IXQsTAzCZ3GqWDc8pfriDGDeg6Sx1OO5GdsPN/5whRM4U44sR\ncqldVoNyS/sKrsLaJbdWjG1aJb5nuFSHzMY8IHms1I34rTsIE5zAPv/DIiYJjJE22r3xW5\n5OOLpRHl5/pWsGClSVKm/1RdtMK4Vbr3S3Ea1acs4QcoeXUAAAADAQABAAAAgCUe+gyTJZ\nltouyqoGgzcYJ8q6EKu+l6wK9sXykmWu5iTSAhEsSDOTG8kKtgmUDfADRk3/AHwZxIxJ69\nuv8L7JcYzJcLoBbo0naICkHNgroHNUqD5APG9+OvSIrM1LMx610cmRle3IxaVdemgRNGLS\n3I0h5MCVBYYXCJnh8dV+ytAAAAQQC94EFjIU/77YA5XtFStRpMaViOunPQaW7MqEgP+fHt\nQL5ZQ86w+Bvn41fJMNtFyaIq4SMhNH3pScBfv7QgubK1AAAAQQDR6t2vthsBw1q2+utFug\nsFVJvEjFZjZXpbCenmm8Vg8INGJnZcygeY86qPhpBvUhmuRX4pKbCQcsU7Hw6/aEbvAAAA\nQQDO8AaF5az9Z38tE5KddwNCSv6OgL5kiNn+ZdpydhAbemthdzww+By9AzQ6TUX26TdNFt\nisgTFay3AqFSEWgOXbAAAABm5vbmFtZQECAwQ=\n-----END OPENSSH PRIVATE KEY-----")
	_, err := LoadPublicKey()

	got := err.Error()

	want := "structure error: tags don't match"

	if !strings.Contains(got, want) {
		t.Errorf("got %s, wanted %s", got, want)
	}
}

func TestPublicKey(t *testing.T) {
	t.Setenv("PUBLIC_KEY", "-----BEGIN PUBLIC KEY-----\nMFkwEwYHKoZIzj0CAQYIKoZIzj0DAQcDQgAEBBVqOkBDYFlsgv8yRp+w79TMhm/I\niCjwbmPk4g75wKdXpCrKvN/myjfjFip4j8lVuOoi9nVrghbCejVguO3PCA==\n-----END PUBLIC KEY-----")

	pk, _ := LoadPublicKey()

	if pk == nil {
		t.Error("public key not loaded")
	}
}
