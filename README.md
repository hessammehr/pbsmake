pbsmake
=======

`pbsmake` is a toy utility written in Go to make it easier to submit jobs to PBS-based batch systems. `pbsmake` looks for all files of a given extension (.com by default since I use it submit Gaussian jobs) and for each it makes a folder containing the input file and a PBS submission script.

Quickstart
----------

Running `pbsmake -help` will give a list of all possible options. Here is my typical workflow.

```sh
# Assuming template.pbs contains your PBS template
$ vim job1.com # Create your job file(s)
$ ./pbsmake
$ cd job1 # For each generated folder run qsub submit-xxx.pbs from inside that folder.
$ qsub submit-job1.pbs
```

The included example `template.pbs` has only been tested on the TORQUE installation in Westgrid for running Gaussian jobs but should work on most similar setups.

I hope you find `pbsmake` useful; pull requests are more than welcome.

License
-------

MIT
