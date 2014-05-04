#!/bin/bash
gfortran lib.f95 example_network.f95 -o example_network
./example_network
rm example_network lib.mod