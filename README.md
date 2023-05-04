# Task 6
Solution for training Tinkoff task №6 in golang
## Task
| Time limit  | Memory limit|
| ----------- | ----------- |
| 1 second    | 256 MB      |

In physical education, there is a division into two teams a<sub>i</sub> . The guys are lined up, each of them has his own height. 
The division into teams will take place according to the “even-odd” principle - all students with even height go to one team, and odd ones go to another.

Unlike the usual lesson, the guys did not line up in height. Instead of the usual order, they got up by chance.
Now physical instructor Yasha looks at the line and thinks - can exactly one pair of students change places so that the teams are the same as in the "first-second" principle. 
In other words, he wants to obtain such an order in which all students with even height are in even positions, and those with odd height are in odd positions.

Help Yasha find the right replacement.

![Task](/task/task.png "Task")
## Input and output data
### Input data format

The first line contains the number *n(2 ≤ n ≤ 1000)* is the number of students in the line.

The next line contains *n* natural numbers *a<sub>i</sub>(1 ≤ a<sub>i</sub> ≤ 10^9)*  - the growth of students.

### Output format

In a single line print *i* and *j* - the numbers of the elements that need to be swapped in order to achieve the given condition *(1 ≤ i, j ≤ n, i != j)*. If there are several answers, any one is allowed.

If there is no way to swap two elements, print - -1, -1

![Input and output data](/task/input_out.png "Input and output data")

## Example

#### Remarks

In the first example, at least one student with an even height will be in an odd position. In the second test, the replacement will result in an incorrect state.

In the third test from the condition, the replacement will bring the line to a valid state [1,2].

### Data example

| Example № | Input | Output |
| ----------- | ----------- | ----------- |
| 1    | n = 4, a<sub>i</sub> = 2 1 4 6 | -1 -1 |
| 2    | n = 2, a<sub>i</sub> = 1 2 | -1 -1 |
| 3    | n = 2, a<sub>i</sub> = 2 1 | 1 2 |

![Example](/task/example.png "Example")
