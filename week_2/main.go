package main 

import (
"fmt"
"os"
"encoding/hex"
"crypto/sha256"
 )

func main (){
	wordsOfPower :="Here lies the secret to ultimate power"
	
	done :=turnToHash (wordsOfPower ) 
	fmt.Printf("This is it %s " ,done)
 
	
	err := makeFile(done,wordsOfPower)
	if err != nil {
    fmt.Println("failed:", err)
    os.Exit(1)
	}
 
 }

func turnToHash (data string) (string){
	result := sha256.Sum256([]byte(data))

	return hex.EncodeToString(result[:])

}
func makeFile(filename string, data string) (error) {

	 err:= os.MkdirAll("blobs/sha256",0755)
	if err != nil{
		return fmt.Errorf("Mkdir error: %w", err) 
	}
	err = os.WriteFile("blobs/sha256/"+filename, []byte(data),0644)
		if err != nil{
		return fmt.Errorf("Write error: %w", err) 
	}

	return nil
}
