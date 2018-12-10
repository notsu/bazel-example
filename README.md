# Bazel with Mono-repo

This is a proof of concept about Bazel and Mono-repositories which services implemented by Golang.

The idea of mono-repo give us many benefits but come with a challenge about how to testing and build fast and efficient. Bazel is the building tools from Google which come with a concept of incremental build and test. So, we can leverage from this concept to build and test fast and solve the challenge of mono-repo.

## Get Started

### Prerequisite

You need to install [Bazel](https://docs.bazel.build/versions/master/install.html) -- a building tool from Google

### Run

- Run `bazel run //:gazelle` to generate `BUILD` for each package including update dependencies
- Run `bazel test //...` to test all service. You can see which test using a cache or not. And it's also testing other services if they depend on each other
- Run `bazel build //...` to build all service. Same as the testing, it uses the cache and builds all dependencies service

## Built with

- [Bazel](https://bazel.build/)
- [Golang](https://golang.org/)

## Author

- Pichet Itngam

## License

This project is licensed under the MIT License