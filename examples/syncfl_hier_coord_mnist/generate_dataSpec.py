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

import subprocess

trainer_nums = [1, 10, 100, 1000, 10000, 100000, 1000000]
dataset_id = "6435345dcb6b63f04a927986"

for trainer_num in trainer_nums:
    dataspec_f = open('dataSpec.json')
    dataspec_g = open(f'dataSpec_{trainer_num}trainers.json', 'w')

    dataspec_f_lines = dataspec_f.readlines()
    dataspec_g.write(dataspec_f_lines[0])
    dataspec_g.write(dataspec_f_lines[1])
    dataspec_g.write(dataspec_f_lines[2])
    dataspec_g.write(dataspec_f_lines[3])
    dataspec_g.write(dataspec_f_lines[4])

    for i in range(trainer_num - 1):
        dataspec_g.write(f"                    \"{dataset_id}\",\n")
    dataspec_g.write(f"                    \"{dataset_id}\"\n")

    dataspec_g.write(dataspec_f_lines[8])
    dataspec_g.write(dataspec_f_lines[9])
    dataspec_g.write(dataspec_f_lines[10])
    dataspec_g.write(dataspec_f_lines[11])
    dataspec_g.close()

    json_f = open('job.json')
    json_g = open(f'job_{trainer_num}trainers.json', 'w')

    json_f_lines = json_f.readlines()
    json_g.write(json_f_lines[0])
    json_g.write(json_f_lines[1])
    json_g.write(json_f_lines[2])
    json_g.write(json_f_lines[3])
    json_g.write(json_f_lines[4])
    json_g.write(json_f_lines[5])
    json_g.write(json_f_lines[6])
    json_g.write(f"    \"dataSpecPath\": \"dataSpec_{trainer_num}trainers.json\",\n")
    json_g.write(json_f_lines[9])
    json_g.write(json_f_lines[10])
    json_g.close()
