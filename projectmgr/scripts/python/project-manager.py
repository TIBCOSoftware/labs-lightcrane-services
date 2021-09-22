#!/usr/bin/env python3

import argparse

def parse_args():
    """Parse command line arguments.
    """
    arg_parse = argparse.ArgumentParser(
        formatter_class=argparse.ArgumentDefaultsHelpFormatter)
    arg_parse.add_argument('-f', '--yml_file', default=None,
                           help='Optional config file for list of services'
                           ' to include.\
                           Eg: python3.6 eis_builder.py -f\
                           video-streaming.yml')
    arg_parse.add_argument('-v', '--video_pipeline_instances', default=1,
                           help='Optional number of video pipeline '
                                'instances to be created.\
                           Eg: python3.6 eis_builder.py -v 6')
    arg_parse.add_argument('-d', '--override_directory',
                           default=None,
                           help='Optional directory consisting of '
                           'of benchmarking configs to be present in'
                           'each app directory.\
                           Eg: python3.6 eis_builder.py -d benchmarking')
    return arg_parse.parse_args()


if __name__ == '__main__':

    # Parse command line arguments
    args = parse_args()

    # Setting number of multi instances variable
    if int(args.video_pipeline_instances) > 1:
        num_multi_instances = int(args.video_pipeline_instances)

    # Start yaml parser
    yaml_parser(args)
