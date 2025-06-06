# My Notes on the Go Programming Language

# Managing Windows Defender Exclusions

When building Go applications, Windows Defender may flag the temporary files created during the build process as potentially unwanted software (PUA).

<img src="images/1749200008499.png" alt="Windows Defender PUA Warning" width="800" />

<img src="images/1749200228340.png" width="500" />

To work around this issue, you can set the `TEMP` and `TMP` environment variables to a different directory within the active session:

```powershell
$env:TEMP = "C:\Users\gregt\AppData\Local\Temp\gotemp"
$env:TMP = "C:\Users\gregt\AppData\Local\Temp\gotemp"
```

Then, in Windows Defender, add an exclusion for the `gotemp` directory:

<img src="images/1749200434394.png" width="500" />
