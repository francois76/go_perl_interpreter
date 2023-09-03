#!/usr/bin/perl
use strict;
use warnings;

my $result = "";
for my $i (0..999) {
    $result .= $i;
}
print $result . "\n";

1;