# Copy folder d1 into folders d2..d25

for i in {2..25}
do
  cp -r d1 d$i
done