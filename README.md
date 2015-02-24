pbsmake
=======

`pbsmake` is a toy utility written in Go to make it easier to submit jobs to PBS-based batch systems. `pbsmake` looks for all files of a given extension (.com by default since I use it to submit Gaussian jobs) and for each it makes a folder containing the original input file and a PBS submission script.

Dependencies
------------
`pbsmake` as a binary has no dependencies. It has been tested on Linux, Mac, Windows, FreeBSD, and OpenBSD (only manually, though).

Quickstart
----------

Running `pbsmake -help` will give a list of all possible options. Here is my typical workflow.

```sh
# Assuming template.pbs contains your PBS template
$ vim job1.com # Create your job file(s)
$ ./pbsmake
$ ls # Your job file has been moved to its own folder
job1  pbsmake
$ cd job1 # For each generated folder run qsub submit-xxx.pbs from inside that folder.
$ qsub submit-job1.pbs
```

The included example `template.pbs` has only been tested on the Westgrid's TORQUE installation with Gaussian jobs but should be easy to adapt to any similar system.

I hope you find `pbsmake` useful; pull requests are more than welcome!

License
-------

MIT
