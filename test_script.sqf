#define freeExtension comment

// Test script "test_script.sqf" to demonstrate 
// the functionality of "callExtension" executable
// Name of the script could be any
// Instructions processed 1 per line
// Lines starting with # are ignored
// Wait for 3.7 seconds...
sleep 1;

// Show extension version
// (this command format is not currently available in Arma 3)
"armago" callExtension "getkeys";

exit;