# Profile Parser

## Description
This go module is auxiliary and serves to link [onnx profiler](https://onnxruntime.ai/docs/performance/tune-performance.html) and [visualizer](https://github.com/XellarReps/visualizer). 
It analyzes the data of two utilities and compares them with each other.

### Compilation and launch

In the root of the project, run:

```bash
go build
./profile_parser *args*
```

### Command line arguments
The module supports different startup modes and support for command line arguments has been added for this.
To find out information about all the arguments, you need to type in the command line after compilation:
```bash
./profile_parser --help
```

Detailed description of each argument:
```text
Usage of ./profile_parser:
  -input_json_path string
        input path (json file with nodes info)
  -input_txt_path string
        input path (txt file with nodes)
  -output_path string
        output path (csv file)
```

### Necessary conditions for the launch
To run, you need to have two files:
1) the json file received when launching the Inference Session (onnx python module) with the setting 
```python
enable_profiling = True;
```
2) the txt file received when running the visualizer with the parameters 
```bash
--info_mode=nodes --write_mode=file
```

It is necessary to specify the paths to these files at startup, and also specify where to put the resulting csv file.