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

if len(sys.argv) != 2:
    raise Exception("Wrong number of arguments; expected 2\n")

total_trainer_num = int(sys.argv[1])

# generate json files
job_id = "622a358619ab59012eabeefb"
task_id = 0
rounds = 5


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


def generate_trainer_config(filename, num):
    # load json data
    input_filename = "config/train_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = f"{str(task_id+num)}"
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()


# make directory for output files
os.system("mkdir output")
os.system("mkdir output/aggregator")
os.system("mkdir output/trainer")


for j in range(total_trainer_num):
    filename = f"config/trainer{j}.json"
    generate_trainer_config(filename, j)

    os.chdir("trainer")
    os.system(
        f"python keras/main.py ../config/trainer{j}.json > "
        + f"../output/trainer/trainer{j}.log 2>&1 &"
    )
    os.chdir("..")

filename = "config/aggregator.json"
generate_agg_config(filename)

os.chdir("aggregator")
os.system(
    "python keras/main.py ../config/aggregator.json > ../output/aggregator/aggregator.log 2>&1 &"
)
os.chdir("..")
