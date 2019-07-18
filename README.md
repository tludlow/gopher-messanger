
# gopher-messanger
A pub/sub client/server for learning more about Go networking abilities


### Future Improvements

 - Add more server-client communication protocols, e.g HTTP, UDP, WebSockets, etc...
 - Evaluate the benefits of having the client receive all published data, and the client deciding if its relevant to them or by having the server only send the data to the client that wants it (who are subscribed)
 - Add some long term data storage methods for the client, e.g cache (probably redis), a database, just storing to a file.
 - Improve upon the message format, including more in depth / data rich messaging between clients.
 - I will surely think of some more interesting examples....
