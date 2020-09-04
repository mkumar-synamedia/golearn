def main():
    num = 1000000
    metric = "cos_stord_processing_duration_seconds"
    help = 'Duration of processing, dimensioned by process and stage'
    type = 'summary'
    sample = 'sum {process = "service-transition-enabled", stage = "gossiper"} 0.027721451'

    f = open("samples.metrics", "w")

    for i in range(num):
        f.write('# HELP ' + metric + '_' + str(i) + ' ' + help + "\n")
        f.write('# TYPE ' + metric + '_' + str(i) + ' ' + type + "\n")
        f.write(metric + '_' + str(i) + '_' + sample + "\n")

    f.close()


if __name__=="__main__":
    main()


# # HELP cos_stord_processing_duration_seconds Duration of processing, dimensioned by process and stage
# # TYPE cos_stord_processing_duration_seconds summary
# cos_stord_processing_duration_seconds_sum {process = "service-transition-enabled", stage = "gossiper"} 0.027721451


