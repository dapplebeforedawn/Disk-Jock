## Shell
```bash
go run disk-jock.go -wav=wavs.dat
go run disk-jock.go -fft=wavs.fft
gnuplot -e "plot 'wavs.fft with lines' ; pause -1"
go run fft.go wavs.dat > wavs-fft.fft
gnuplot -e "plot 'wavs.fft' with lines; pause -1"

go run peak-detect.go kawehi-mj-mag-only.fft
gnuplot -e "plot 'kawehi-mj.fft' with lines, 'kawehi-mj-peaks.dat' ls 3; pause -1"
```

## VIM
```vim
%s/^/\=printf('%-4d ', 43*line('.'))
%s/\(\d\+\)/\=(submatch(1)*43)/
%s/$/ 0/
```
