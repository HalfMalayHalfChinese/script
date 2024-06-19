# Check if the directory exists, if not, create it
if (-Not (Test-Path -Path "C:\temp\testDir" -PathType Container)) {
    New-Item -ItemType Directory -Path "C:\temp\testDir"
}

# Change to the directory
Set-Location -Path "C:\temp\testDir"

# Create an empty file
New-Item -ItemType File -Name "test.ps1"

# Capture network configuration and save to a log file
Get-NetIPConfiguration > ifconfig.log
