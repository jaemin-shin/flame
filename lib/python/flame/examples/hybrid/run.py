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
if len(sys.argv) < 4 or len(sys.argv) != int(sys.argv[2]) + 3:
    # logger?
    raise Exception(
        "Wrong number of arguments; \nExample usage: python run.py"
        + "(num of total trainers) (num of clusters) (cluster 0 trainers) (cluster 1 trainers) ... (cluster n trainers)"
    )

total_trainer_num = int(sys.argv[1])
total_cluster_num = int(sys.argv[2])

cluster_trainer_num = []
for i in range(total_cluster_num):
    cluster_trainer_num.append(int(sys.argv[3 + i]))

if sum(cluster_trainer_num) != total_trainer_num:
    raise Exception(
        "The number of trainers for each clusters do not sum up to the trainer num"
    )

# generate json files
job_id = "622a358619ab59012eabeefb"
task_id = 0
rounds = 50


def generate_agg_config(filename):
    # load json data
    input_filename = "config/agg_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = str(task_id)
    data["channels"][0]["backend"] = "mqtt"
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()


def generate_trainer_config(filename, num, cluster_num, cluster_value):
    # load json data
    input_filename = "config/train_template.json"
    file = open(input_filename, "r")
    data = json.load(file)
    data["job"]["id"] = job_id
    data["taskid"] = f"{str(task_id+num+(cluster_num+1)*1000)}"
    data["groupAssociation"]["param-channel"] = "cluster_" + str(cluster_num)

    data["channels"][0]["groupBy"]["value"] = cluster_value
    data["channels"][0]["backend"] = "p2p"
    data["channels"][1]["backend"] = "mqtt"
    data["hyperparameters"]["rounds"] = rounds

    # save json data
    file = open(filename, "w+")
    file.write(json.dumps(data, indent=4))
    file.close()


# make directory for output files
os.system("mkdir output")
os.system("mkdir output/aggregator")
os.system("mkdir output/trainer")

cluster_value = []
for i in range(total_cluster_num):
    cluster_value.append("cluster_" + str(i))

for i in range(total_cluster_num):
    for j in range(cluster_trainer_num[i]):
        filename = f"config/cluster{i}_trainer{j}.json"
        generate_trainer_config(filename, j, i, cluster_value)

        os.chdir("trainer")
        os.system(
            f"python main.py ../config/cluster{i}_trainer{j}.json > "
            + f"../output/trainer/cluster{i}_trainer{j}.log 2>&1 &"
        )
        os.chdir("..")

filename = "config/aggregator.json"
generate_agg_config(filename)

os.chdir("aggregator")
os.system(
    "python main.py ../config/aggregator.json > ../output/aggregator/aggregator.log 2>&1 &"
)
os.chdir("..")
