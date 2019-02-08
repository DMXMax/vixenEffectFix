package main

import (
    "bufio"
    "flag"
    "fmt"
    "log"
    "os"
    "io"
    "encoding/binary"
  )

  func main(){
    fptr := flag.String("infile", "input.eseq", "file path to read")
    optr := flag.String("outfile", "out.eseq", "file path to write")
    flag.Parse()
    
    //Open the Input file, report an error if needed 
    f,err := os.Open(*fptr)

    if err != nil{
      log.Fatal(err)
    }
    //Open the Output file, report an error. This will truncate an
    //existing file
    out,err := os.Create(*optr)
    if err != nil{
      log.Fatal(err)
    }

    w := bufio.NewWriter(out)
    //Take advantage of defer to just clean up. 
    defer func(){
      if err = f.Close(); err != nil{
        log.Fatal(err)
      }
      if err = w.Flush(); err != nil {
        log.Fatal(err)
      }
      
      if err = out.Close(); err != nil{
        log.Fatal(err)
      }
    }()

    var bcount int64 = 0;
    r := bufio.NewReader(f)
    //A buffer size of 4096 works well
    b := make([]byte, 4096)
    for {
      k, err := r.Read(b)
    // This line checks if we're at the start of the file, where we
    // make our change at byte 12. We'll write an int32 of 1 to change the
    // start model_start. 
      if bcount == 0 {
        binary.LittleEndian.PutUint32(b[12:16], uint32(1))
      }
      bcount += int64(k)
      if err != nil{
        if err != io.EOF{
        fmt.Println("Error reading file:", err)
      }
        fmt.Println()
        break
      }
      _, err = w.Write(b[:k])
      if err != nil{
        fmt.Println("Error Writing File:", err)
      }
    }
    fmt.Printf("%v bytes written to %v.\n", bcount, *optr) 
  }
