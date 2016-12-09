
docker run -it --entrypoint bash golang:1.8beta1-wheezy

You should now be at the command prompt inside your container:
root@cc5dae79d6fe:/go#
Then run:
go get github.com/jeremywho/pluginTest && cd $GOPATH/src/github.com/jeremywho/pluginTest
You should now be ready to follow along!

Step 1 - Create the plugin

Our plugin:
package main

func Add(x, y int) int {
    return x+y
}
A simple method to add to integers. That’s it for this example.
Step 2 - Build the plugin

According to the documentation we can build plugins with:
go build -buildmode=plugin
We are going to be a bit more verbose with our build command:
go build -buildmode=plugin -o myplugin.so myplugin.go
Let’s specify the output file with -o myplugin.so and the input file at the end of the command
If you take a look in your directory, you should see the myplugin.so library:
root@cc5dae79d6fe:/go/src/github.com/jeremywho/pluginTest# ls
main.go  myplugin.go  myplugin.so
Now that we have a shared library, we can…
Step 3 - Load the plugin

Make sure the plugin package is imported.
Load the shared library file we built:
p, _ := plugin.Open("./myplugin.so")  
We call plugin.Open() with the path to the shared library file we created in the previous step. This gives us back a pointer to a Plugin type.
Next we call plugin.Lookup() with the name of the Symbol we want to use:
add, _ := p.Lookup("Add")
A Symbol can be “any exported variable or function”.
Step 4 - Use it!

Here is how we call your plugin method:
sum := add.(func(int, int) int)(1, 2)
This was confusing to me at first, let’s break it down.
The original plugin function signature looks like:
func Add(x, y int) int
Which is a function that takes two integers func(int,int) and returns an integer:
func(int,int) int
We do a type assertion (think cast) on our Symbol add
add.(func(int, int) int)
which tells Go that add is a function with that signature.
Lastly we call it with the parameters we want to add (1,2) and store the results in sum
sum := add.(func(int, int) int)(1, 2)
Thats it! We can run our program and see the results:
root@cc5dae79d6fe:/go/src/github.com/jeremywho/pluginTest# go run main.go
3    
