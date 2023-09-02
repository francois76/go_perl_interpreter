
        use strict;
        use warnings;
        use JSON qw(from_json to_json);


        sub print_ad15d99e0a9045bf86a483dc879bc70b
        {
                print '[ad15d99e0a9045bf86a483dc879bc70b - PRINT]', @_;
        }


        sub main
        {
                my $firstKey = from_json("{\"AString\":\"test\",\"AInt\":52,\"ABool\":true}");
my $secondKey = from_json("2");

print_ad15d99e0a9045bf86a483dc879bc70b "$firstKey->{AString} \n" ,"toto";
print_ad15d99e0a9045bf86a483dc879bc70b "$firstKey->{AInt} \n";
print_ad15d99e0a9045bf86a483dc879bc70b "$firstKey->{ABool} \n";
print_ad15d99e0a9045bf86a483dc879bc70b "$secondKey\n";
return;

        }
        my $result = main();
        print to_json($result);
        1;