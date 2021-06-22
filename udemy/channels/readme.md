```golang
// send the value '5' into the channel
channel <- 5
// Wait for a value to be sent into the channel.  When we get one, assign the value to 'myNumber'
myNumber <- channel
// or simply print it out right away
fmt.Println(<- channel)
```