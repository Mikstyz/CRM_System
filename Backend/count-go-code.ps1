[Console]::OutputEncoding = [System.Text.Encoding]::UTF8

$totalLines = 0
$totalChars = 0

$files = Get-ChildItem -Path . -Filter *.go -Recurse -File -ErrorAction SilentlyContinue

foreach ($file in $files) {
    if (-not (Test-Path $file.FullName)) {
        continue
    }

    $lines = Get-Content $file.FullName -ErrorAction SilentlyContinue
    foreach ($line in $lines) {
        $trimmed = $line.Trim()
        if ($trimmed -ne "") {
            $totalLines++
            $noWhitespace = $trimmed -replace "[ \t]", ""
            $totalChars += $noWhitespace.Length
        }
    }
}

Write-Output "`n==========================="
Write-Output "✅ Total lines of code: $totalLines"
Write-Output "✅ Total characters (no spaces/tabs): $totalChars"
Write-Output "==========================="

# Keep the window open
Read-Host -Prompt "`nPress Enter to close this window"
