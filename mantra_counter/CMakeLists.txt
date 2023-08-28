cmake_minimum_required(VERSION 3.25)
project(counter)


find_package(wxWidgets REQUIRED COMPONENTS core base)
include(${wxWidgets_USE_FILE})

add_executable(counter
        src/main.cpp)


target_link_libraries(counter ${wxWidgets_LIBRARIES})


set_target_properties(counter PROPERTIES
        CXX_STANDARD 11
        CXX_STANDARD_REQUIRED ON
)
