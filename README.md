# Snowflake-Golang

SnowFlake-Golang is a golang implemtentaion of twitter's snowflake-jave-uniqueID generation algorithm.
It is simple sequence generator that generates 64-bit IDs based on the concepts outlined in the Twitter snowflake service

According to [twitter-snowflake](https://github.com/twitter-archive/snowflake/tree/snowflake-2010):

>Snowflake is a network service for generating unique ID numbers at high scale with some simple guarantees.

[![Build Status](https://travis-ci.com/RO-29/snowflake-Golang.svg?branch=master)](https://travis-ci.com/RO-29/snowflake-Golang)

## Requirements

### Performance

 1 minimum 10k ids per second per process
 2 response rate 2ms (plus network latency)

### Uncoordinated

For high availability within and across data centers, machines generating ids should not have to coordinate with each other.

### (Roughly) Time Ordered

We have a number of API resources that assume an ordering (they let you look things up "since this id").

However, as a result of a large number of asynchronous operations, we already don't guarantee in-order delivery.

We can guarantee, however, that the id numbers will be k-sorted (references: http://portal.acm.org/citation.cfm?id=70413.70419 and http://portal.acm.org/citation.cfm?id=110778.110783) within a reasonable bound (we're promising 1s, but shooting for 10's of ms).

### Directly Sortable

The ids should be sortable without loading the full objects that the represent. This sorting should be the above ordering.

### Compact

There are many otherwise reasonable solutions to this problem that require 128bit numbers. For various reasons, we need to keep our ids under 64bits.

##  Solution

* id is composed of:
  * time - Epoch timestamp in milliseconds precision - 42 bits. The maximum timestamp that can be represented using 42 bits is 242 - 1, or 4398046511103, which comes out to be Wednesday, May 15, 2109 7:35:11.103 AM. That gives us 139 years with respect to a custom epoch.
  * configured machine id - 10 bits. This gives us 1024 nodes/machines.
  * sequence number - 12 bits - rolls over every 4096 per machine (with protection to avoid rollover in the same ms)

Your microservices can use this Sequence Generator to generate IDs independently. This is efficient and fits in the size of a bigint.

### System Clock Dependency

You should use NTP to keep your system clock accurate.  Snowflake protects from non-monotonic clocks, i.e. clocks that run backwards.  If your clock is running fast and NTP tells it to repeat a few milliseconds, snowflake will refuse to generate ids until a time that is after the last time we generated an id. Even better, run in a mode where ntp won't move the clock backwards. See http://wiki.dovecot.org/TimeMovedBackwards#Time_synchronization for tips on how to do this.

## Install

It requires the latest stable Go version.

`make`

Executables are compiled in the `build` directory.

It can be imported as a package in your project, this project follows golang modules and semver approach

## Run

Run the executables.

Flags documentation is available with `-h`.
