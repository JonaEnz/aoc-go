param($day, [String]$year = "2022")

if (!(Test-Path "./.session")) {
  throw "No /.session found"
}

$sessionCode = Get-Content ".\.session"

if ($null -eq $day) {
  $day = (Get-Date).Day
}
  
$downloadToPath = ".\$year\day$day\input.txt"
if (!(Test-Path ".\$year\day$day")) {
  New-Item -Path "." -name "$year\day$day" -ItemType "directory" 
}
$remoteFileLocation = "https://adventofcode.com/$year/day/$day/input"
  
$session = New-Object Microsoft.PowerShell.Commands.WebRequestSession
    
$cookie = New-Object System.Net.Cookie 
    
$cookie.Name = "session"
$cookie.Value = $sessionCode
$cookie.Domain = ".adventofcode.com"

$session.Cookies.Add($cookie);

Invoke-WebRequest $remoteFileLocation -WebSession $session -TimeoutSec 900 -OutFile $downloadToPath

Copy-Item ".\template\template.go" -Destination ".\$year\day$day\day$day.go"

Start-Process "https://adventofcode.com/$year/day/$day"