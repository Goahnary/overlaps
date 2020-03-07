# overlaps
Tired of going to the grocery and only catching one of the 3 sales they have every week?? Want an unnecessarily complex solution? Well you're in luck!

Say you have some sales events at your grocery store:

```
            day   0  1  2  3  4  5     
apple sale        [-----------]
banana sale             [--------]
pickle sale                [-------------]
```

When will these events overlap?

Well if you run this program after modifying the data in the main function to represent your sales events... You can find out!

run the program with:
```
go run overlaps.go
```

your output should look something like this:

```
-------------
| Events    |
-------------

Name: apple sale, 
Start:0
End:4

Name: banna sale, 
Start:2
End:5

Name: pickle sale, 
Start:3
End:10


-------------
| Overlaps  |
-------------

Events: pickle sale, apple sale, 
Start:3
End:4

Events: pickle sale, banna sale, 
Start:3
End:5

Events: banna sale, apple sale, 
Start:2
End:4
```

Then you're done! You now know when all the sales will be on the same day and you can make that sweet apple, banana, pickle sandwich!
