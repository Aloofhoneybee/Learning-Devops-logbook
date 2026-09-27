**<ins>Week 2</ins>**

This week I worked on trying to understand containers a little bit more. I now understand that containers technology like docker is basically an abstraction of the capabilities already present in the Linux system namely namespaces, cgroups, veth and others to emulate a file system.
I'm also trying to use arch more and typed this using neo vim to try get more used to it.

<img width="1280" height="960" alt="book 1" src="https://github.com/user-attachments/assets/1c9fb173-1787-4de4-b5f6-672febd83fd0" />
[Trying to understand more about container images and how the oci model affects it]

<img width="1280" height="960" alt="book 2" src="https://github.com/user-attachments/assets/d592df29-9c96-4857-92f7-49aecb6fbb9b" />

![Trying to understand the difference between image manifest and index manifest]

(forgive my horrible handwriting) 

Then I tried to understand the OCI image spec at a higher level what I just understood initially that it was setup by Docker and the Linux Foundation to have a standard as to How images should be built but now I understand more about it's content is at a higher level: what a manifest is, what  a config is, a digest, a index manifest and what a descriptor is.

So now I'm going to move on to the Runtime spec and work on understanding it along side config.json and runc.
