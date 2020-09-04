
use strict;
use Benchmark qw(:all) ;

sub read_lines {
    my ($file) = @_;
 
    open my $in, "<", $file;
    local $/ = undef;
    my $content = <$in>;
    close $in;
    return split /\n/, $content;
}

my @lines = read_lines('samples.metrics');

my $help = qr/^(#\s+HELP\s+([^\s]+).*)$/;
my $type_metric = qr/^(#\s+TYPE\s+([^\s]+).*)$/;
my $sample = qr/^([^\#][^\s|{]*)\s*({(.*)})?\s+([^\s]+)$/;

sub process { 
	foreach my $line (@lines) {
	   if ($line =~ m/$help/) # ~m is the match operator
	   {

	   }
	   elsif ($line =~ m/$type_metric/)
	   {
	      
	   }
	   elsif ($line =~ m/$sample/)
	   {
	     
	   }
	}
}



my $result = undef;





$result = timethis (10, 'process');

print "total CPU = ", $result->cpu_a, "\n";








