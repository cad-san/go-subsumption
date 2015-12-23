# go-subsumption

[![Build Status](https://travis-ci.org/cad-san/go-subsumption.svg?branch=master)](https://travis-ci.org/cad-san/go-subsumption)
[![Coverage Status](https://coveralls.io/repos/cad-san/go-subsumption/badge.svg?branch=master&service=github)](https://coveralls.io/github/cad-san/go-subsumption?branch=master)

Implementation of "Subsumption Architecture" with Golang

##Overview
This is simple library to implements "subsumption architecture" devised by Rodney Allen Brooks.

This library doesn't implement actuators and sensors, but we provide the interfaces for sensing the environment and behavior as reactions.

So you can make the concrete implements weather physical robot or simulator.
