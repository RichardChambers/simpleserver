package main

import  (
        "fmt"
        "log"
        "net/http"
)

var gcno=1

func main () {
    gcno = 1
    // start up the server to accept HTTP transmissions and
    // when one comes in, hand it off to the handler function.
    http.HandleFunc("/", handler)
    log.Fatal(http.ListenAndServe("192.168.0.5:8282", nil))
}

func handler (w http.ResponseWriter, r *http.Request) {
    // all for cross domain scripting so that a browser can start up
    // a scriptted HTML page and the JavaScript http request will work.
    w.Header().Set("Access-Control-Allow-Origin", "*")
    w.Header().Set("Access-Control-Allow-Methods", "POST GET")

    // just for grins we are going to ignore every fourth request
    // by returning a gcno of zero to indicate no pending transactions
    // waiting. this helps to exercise the JavaScript for such an event.
    if (gcno % 4 != 0) {
      // generate a transaction with transaction data.
      fmt.Fprintf (w, "{ \"gcno\":%d, \"trans\": [", gcno)
      fmt.Fprintf (w, "{\"quant\":1,\"mnemonic\":\"hamburger\"},")
      fmt.Fprintf (w, "{\"quant\":2,\"mnemonic\":\"hamburg_2\"},")
      fmt.Fprintf (w, "{\"quant\":3,\"mnemonic\":\"hamburg_3\"},")
      if (gcno % 7 == 0) {
        // every now and then add some additional items to the transaction.
        fmt.Fprintf (w, "{\"quant\":2,\"mnemonic\":\"fries_1\"},")
        fmt.Fprintf (w, "{\"quant\":3,\"mnemonic\":\"fries_2\"},")
        fmt.Fprintf (w, "{\"quant\":4,\"mnemonic\":\"fries_3\"},")
      }
      fmt.Fprintf (w, "{\"quant\":4,\"mnemonic\":\"hamburg_4\"}")
      fmt.Fprintf (w, "] }");
    } else {
      // generated a No Transaction Available message.
      fmt.Fprintf (w, "{ \"gcno\":0, \"trans\": [")
      fmt.Fprintf (w, "{\"quant\":0,\"mnemonic\":\"hamburg_4\"}")
      fmt.Fprintf (w, "] }");
    }
    gcno++
}
