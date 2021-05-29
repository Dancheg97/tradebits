package main

import (
	"fmt"
	"os"
	"os/exec"
)

func main() {
	fmt.Println(
		"\033[36m",
		`
      _____                    _____                    _____                _____                    _____          
     /\    \                  /\    \                  /\    \              /\    \                  /\    \         
    /::\    \                /::\    \                /::\    \            /::\    \                /::\    \        
    \:::\    \              /::::\    \              /::::\    \           \:::\    \              /::::\    \       
     \:::\    \            /::::::\    \            /::::::\    \           \:::\    \            /::::::\    \      
      \:::\    \          /:::/\:::\    \          /:::/\:::\    \           \:::\    \          /:::/\:::\    \     
       \:::\    \        /:::/__\:::\    \        /:::/__\:::\    \           \:::\    \        /:::/__\:::\    \    
       /::::\    \      /::::\   \:::\    \       \:::\   \:::\    \          /::::\    \       \:::\   \:::\    \   
      /::::::\    \    /::::::\   \:::\    \    ___\:::\   \:::\    \        /::::::\    \    ___\:::\   \:::\    \  
     /:::/\:::\    \  /:::/\:::\   \:::\    \  /\   \:::\   \:::\    \      /:::/\:::\    \  /\   \:::\   \:::\    \ 
    /:::/  \:::\____\/:::/__\:::\   \:::\____\/::\   \:::\   \:::\____\    /:::/  \:::\____\/::\   \:::\   \:::\____\
   /:::/    \::/    /\:::\   \:::\   \::/    /\:::\   \:::\   \::/    /   /:::/    \::/    /\:::\   \:::\   \::/    /
  /:::/    / \/____/  \:::\   \:::\   \/____/  \:::\   \:::\   \/____/   /:::/    / \/____/  \:::\   \:::\   \/____/ 
 /:::/    /            \:::\   \:::\    \       \:::\   \:::\    \      /:::/    /            \:::\   \:::\    \     
/:::/    /              \:::\   \:::\____\       \:::\   \:::\____\    /:::/    /              \:::\   \:::\____\    
\::/    /                \:::\   \::/    /        \:::\  /:::/    /    \::/    /                \:::\  /:::/    /    
 \/____/                  \:::\   \/____/          \:::\/:::/    /      \/____/                  \:::\/:::/    /     
                           \:::\    \               \::::::/    /                                 \::::::/    /      
                            \:::\____\               \::::/    /                                   \::::/    /       
                             \::/    /                \::/    /                                     \::/    /        
                              \/____/                  \/____/                                       \/____/         
                                                                                                                     
`,
		"\033[0m",
	)
	os.RemoveAll("data/base")
   os.RemoveAll("data/data")
   os.RemoveAll("market/data")
   os.RemoveAll("user/data")
   exec.Command("go", "test", "./...")
	os.RemoveAll("data/base")
   os.RemoveAll("data/data")
   os.RemoveAll("market/data")
   os.RemoveAll("user/data")
   fmt.Println("TESTING FINISHED")
}
