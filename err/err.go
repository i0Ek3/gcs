package main

func checkError(err error, info... string) {
    if err != nil {
        if len(info) > 0 {
            panic("Error: " + info[0] + " " + err.Error())
        } else {
            panic("Error: " + err.Error())
        }
    }
}
