# polymerGolangDemo
A quick and dirty demo showing how to use Golang and Polymer together - with templates as well as core-ajax

This is a quick demo to show how to use Golang and Polymer together. There are two basic models:

1. Use Golang solely as a RESTful source to interact with polymer through core-ajax
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

