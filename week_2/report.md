**<ins>Week 2</ins>**

This week I worked on trying to understand containers a little bit more. I now understand that containers technology like docker is basically an abstraction of the capabilities already present in the Linux system namely namespaces, cgroups, veth and others to emulate a file system.
I'm also trying to use arch more and typed this using neo vim to try get more used to it.

![Trying to understand more about container images and how the oci model affects it](Learning-Devops-logbook\images\week 2 img\book 1.jpg)

![Trying to understand the difference between image manifest and index manifest](Learning-Devops-logbook\images\week 2 img\book 2.jpg)

(forgive my horrible handwriting) 

Then I tried to understand the OCI image spec at a higher level what I just understood initially that it was setup by Docker and the Linux Foundation to have a standard as to How images should be built but now I understand more about it's content is at a higher level: what a manifest is, what  a config is, a digest, a index manifest and what a descriptor is.

So now I'm going to move on to the Runtime spec and work on understanding it along side config.json and runc.
