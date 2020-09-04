import re
import timeit

reps = 10


def process(lines, metric_help, metric_type, metric_sample):
    for line in lines:
        if metric_help.match(line):
            pass
        elif metric_type.match(line):
            pass
        elif metric_sample.match(line):
            pass


if __name__ == '__main__':

    f = open("samples.metrics", "r")

    lines = f.read().splitlines()

    metric_help = re.compile('^(#\s+HELP\s+([^\s]+).*)$')
    metric_type = re.compile('^(#\s+TYPE\s+([^\s]+).*)$')
    metric_sample = re.compile('^([^\#][^\s|{]*)\s*({(.*)})?\s+([^\s]+)$')

    start = timeit.default_timer()

    for i in range(reps):
        process(lines, metric_help, metric_type, metric_sample)

    print((timeit.default_timer() - start) / reps)
