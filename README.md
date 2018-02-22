# PromptStatementOne
This repository contains the go source code of a small excecutable to shorten my work directory using elipsis.   

## Get started
1. Compile the binary - `go build pwdname.go`
2. Add the binary to your environment path   
    ```bash
    mkdir -p ~/local/bin/
    echo "export PATH=$PATH:~/local/bin" >> .bash_profile
    ln -s $PWD/pwdname ~/local/bin/pwdname
    ```
3. Include it in your PS1 variable in your `bash_profile`   
    `export PS1="\[\033[38;5;48m\]\@\[\e[m\] \$(pwdname) \[\e[32m\]=>\[\e[m\] "`

## Why
This was previous written using `sed` and `awk`. There wasn't anything wrong with it, I just wanted to write something in go.   
This is my hello-world for go.   

Here's the old one... which is about 3x quicker   
```bash
function generate_pwd {
  pwd | sed s.$HOME.~.g | awk -F"/" '
  BEGIN { ORS="/" }
  END {
  for (i=1; i<= NF; i++) {
      if ((i == 1 && $1 != "") || i == NF-1 || i == NF) {
        print $i
      }
      else if (i == 1 && $1 == "") {
        print "/"$2
        i++
      }
      else {
        print ".."
      }
    }
  }'
}
```