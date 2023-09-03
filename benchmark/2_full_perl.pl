#!/usr/bin/perl
use strict;
use warnings;

my $result = "";
for my $i (0..99999) {
    $result .= $i;
}
print length($result) . "\n";

1;