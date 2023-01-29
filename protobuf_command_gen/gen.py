from os import walk
import argparse
import os
import subprocess

parser = argparse.ArgumentParser()
parser.add_argument('--proto_path')
parser.add_argument('--out_path')
parser.add_argument('--go_package')
parser.add_argument('input', nargs='+')
args = parser.parse_args()

end_cmd = "";

end_cmd += "protoc --experimental_allow_proto3_optional --proto_path={0} --go_out={1} --go_opt=paths=source_relative --go-grpc_out={1} --go-grpc_opt=paths=source_relative".format(args.proto_path, args.out_path)

for inp in args.input:
    dirr = os.path.dirname(inp)
    end_cmd += " --go_opt=M{0}={1}/{2}".format(inp, args.go_package, dirr)
    end_cmd += " --go-grpc_opt=M{0}={1}/{2}".format(inp, args.go_package, dirr)

end_cmd += " ";
end_cmd += " ".join(args.input)

print(end_cmd)
p = subprocess.Popen(end_cmd, shell=True, stdout=subprocess.PIPE, stderr=subprocess.STDOUT)
for line in p.stdout.readlines():
    print(line),
retval = p.wait()
