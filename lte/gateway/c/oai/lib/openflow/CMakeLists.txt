# Copyright (c) 2017-present, Facebook, Inc.
# All rights reserved.

# This source code is licensed under the BSD-style license found in the
# LICENSE file in the root directory of this source tree. An additional grant
# of patent rights can be found in the PATENTS file in the same directory.

set(OUTPUT_DIR "${CMAKE_CURRENT_BINARY_DIR}")

include_directories("${OUTPUT_DIR}")
include_directories("/usr/include/openvswitch")

# add_library(LIB_OPENFLOW
#   # ovs_tunnel_flows.c
#   # ovs_utils.c
# )
# target_link_libraries(LIB_OPENFLOW
#   COMMON
#   LIB_BSTR
# )
# target_include_directories(LIB_OPENFLOW PUBLIC
#   ${CMAKE_CURRENT_SOURCE_DIR}
# )
