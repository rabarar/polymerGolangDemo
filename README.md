# polymerGolangDemo
A quick and dirty demo showing how to use Golang and Polymer together - with templates as well as core-ajax

This is a quick demo to show how to use Golang and Polymer together. There are two basic models:

1. Use Golang solely as a RESTful source to interact with polymer through core-ajax (see http://localhost:8080/feeds/api/rob2)

2. Use Golang templates + Polymer to interact with componts in DOM.

DISCLAMER!

I am not a front-end guy at all - so this demo is not pretty!! It's meant more to show the mechanics of how things work.

Setup

Polymer Setup:
Use bower to install polymer and its components. See the bower.json file to start

Golang Setup:
I'm using gin-gonic for my web framework. See the imports in the webSever.go file for the gin package.

Runtme Setup:
All static content is installed ./static
All templates for the web server are installed in ./templates


Some interesting things to note.
Look at the templates and you'll see {{ }} as well as {[ ]} - the first set of delimiters is for polymer - not Golang templates! In the webSever.go you'll see that I tell golang templates to use the latter pair of delimeters.

Also notice how the linkage between the webServer GET methods and the templates work together. Each GET method has a local context that can be passed to the templates. I demonstrate how one might use this to programmatically build out polymer tags and associated data from golang.


