### Why: Proxy Port?


Assuming you are talking about Uncle Bob's Clean Architecture...
Then third party API calls should be treated just like a database call,
or any other third-party dependency. It belongs in the outermost circle
in this diagram, where things like HTTP libraries absolutely belong.
