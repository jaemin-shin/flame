# Copyright 2023 Cisco Systems, Inc. and its affiliates
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#      http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.
#
# SPDX-License-Identifier: Apache-2.0

import os
import sys
import json

# get trial
if len(sys.argv) != 3:
    # logger?
    raise Exception(
        "Wrong number of arguments; \nExample usage: python run.py"
        + "(num of total trainers) (num of mid aggregators)"
    )

total_trainer_num = int(sys.argv[1])
total_midagg_num = int(sys.argv[2])

# generate json files
job_id = "622a358619ab59012eabeefb"
task_id = 0
rounds = 41


def generate_agg_config(filename):
    # load json data
    input_filename = "config/agg_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = str(task_id)
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()

def generate_coord_config(filename):
    # load json data
    input_filename = "config/coord_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = str(task_id + 1000000)
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()

def generate_mid_agg_config(filename, cluster_num):
    # load json data
    input_filename = "config/mid_agg_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = f"{str(task_id+(cluster_num+1)*1000)}"
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()


def generate_trainer_config(filename, num):
    # load json data
    input_filename = "config/train_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = f"{str(task_id+num+10000)}"
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()


# make directory for output files
os.system("mkdir output")
os.system("mkdir output/coordinator")
os.system("mkdir output/aggregator")
os.system("mkdir output/middle_aggregator")
os.system("mkdir output/trainer")

for i in range(total_trainer_num):
    filename = f"config/trainer{i}.json"
    generate_trainer_config(filename, i)

    os.chdir("trainer")
    os.system(
        f"python main.py ../config/trainer{i}.json > "
        + f"../output/trainer/trainer{i}.log 2>&1 &"
    )
    os.chdir("..")

for i in range(total_midagg_num):
    filename = f"config/mid_aggregator{i}.json"
    generate_mid_agg_config(filename, i)

    os.chdir("middle_aggregator")
    os.system(
        f"python main.py ../config/mid_aggregator{i}.json > "
        + f"../output/middle_aggregator/mid_aggregator{i}.log 2>&1 &"
    )
    os.chdir("..")

filename = "config/aggregator.json"
generate_agg_config(filename)

os.chdir("top_aggregator")
os.system(
    "python main.py ../config/aggregator.json > ../output/aggregator/aggregator.log 2>&1 &"
)
os.chdir("..")

filename = "config/coordinator.json"
generate_coord_config(filename)

os.chdir("coordinator")
os.system(
    "python main.py ../config/coordinator.json > ../output/coordinator/coordinator.log 2>&1 &"
)
os.chdir("..")