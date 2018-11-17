keywords:
- typedef
- class
-

newtype Tree a = Node (a, [Node a]) | Leaf a
newtype Decimal = (val: Integer, point: Integer)

class Equatable a {
    (=) :: (expr1:a, expr2:a) returns Bool
    (!=) :: (expr1:a, expr2:a) returns Bool
}

class Orderable a {
    Equatable

    (<)
    (>)
    (>=)
    (<=)
}

class Number a {
    Equatable, Orderable, Enumerable

    (+) :: (expr1:a, expr2:a) returns result:a
    (-) :: (expr1:a, expr2:a) returns result:a
    (*) :: (expr1:a, expr2:a) returns result:a

    negate :: (expr:a) returns neg:a
    abs :: (expr:a) returns value:a
}



() = type struct = func()
xs = [1,2,3,4,5,6,7,8,9]

filter(predicate:(even), collections:(xs))

a: some value
a -1
if a >= 0
list[a]

-- Derives is haskell magic for auto instantiate.
newtype Integer = -inf | ... | 0 | ... | inf
newtype Int = -MIN_INT | -MIN_INT+1 | ... | 0 | 1 | ... | MAX_INT

-- Pseudo
-- Do class instances need to have the same named parameters.

newtype Negative = Integer a ! a < 0

Integer:Number {
    -- (+), reads as add, computes numerical addition between two integers.
    (+) (expr1:Integer, expr2:Integer) returns Integer {
        return expr1 + expr2
    }

    -- (-), reads as sub, computes numerical subtraction between two integers.
    (-) (expr1:Integer, expr2:Integer) returns result:Integer {
        return expr1 - expr2
    }

    -- (*), reads as mul, computes numerical multiplication between two integers.
    (*) (expr1:Integer, expr2:Integer) returns result:Integer {
        return expr1 * expr2
    }

    -- negate returns the negated value of an integer.
    negate(expr:Integer) returns result:Integer {
        return -expr1
    }

    (collections:xs) contains(element:x) returns Bool
    contains(collections:xs, element:x) returns Bool

    -- abs returns the absolute value of an integer.
    abs(expr:Integer) returns result:Integer {
        if expr < 0 {
            return negate(expr)
        } else {
            return expr
        }
    }
}

type Point = (x: Number, y: Number)

Bool
String

-- TODO(_): Which one?
typedef Bool = True | False
typedef Bool = :true | :false
typedef Bool = true | false

-- TODO(_): Do we enforce named return values?
and :: (expr1:Bool, expr2:Bool) returns Bool
or  :: (expr1:Bool, expr2:Bool) returns Bool
not :: (expr:Bool)              returns Bool

class Comparable a {
    (=) :: (expr1:a, expr2:a) returns Bool
    (!=) :: (expr1:a, expr2:a) returns Bool
}
\x -> x+1
func(x:a) -> x+1

tuple = struct? = func (maybe)

:light = ArmorType iota
:medium
:heavy

struct Armor {
    type ArmorType
    level Integer
}

struct Mob {
    name String
    health Integer
    armor Armor
}

filter (\x -> not $ primes `contains` x) [1..10]


xs = [1, 2, 3, 4]
x = 3
contains(collection:xs, element:x) = contains(element:x, collection:xs)

> true

-- This is impossible?
xs.contains(element: x)
xs `contains` x

-- This will create an anonymous curried function where the
named parameter collections will be filled as a positional argument.

xs.filter(contains(collection:_, element:x))

filter(predicate:even, functor:[1..10])

onlyEven = filter(predicate:even)

-- Galvin
function quicksort(Orderable a => list:[a]) returns sorted:[a] {
    | list:[]         = []   -- match empty list
    | list:[x]        = [x]  -- match list with only one element
    | list:(pivot:xs) = quicksort(list:lesser) . pivot . quicksort(list:greater)
    let lesser  = filter(predicate:(\x -> x <= pivot), collection:xs) -- Both of these work
    let greater = filter(predicate:(\x -> x > pivot))(collection:xs)
}
filter isEven . map (+1) $ [1..3]

filter(predicate:(isEven), collection:(map(verb:(\x -> x+1), collection:([1..3]))))

let oneHigher = map((\x -> x+1), [1..3])
let onlyEven  = filter(isEven, oneHigher)

filter anotherPred . filter isEven . map (\x -> x+1) [1..3]



(filter anotherPred . filter isEven . map (\x -> x+1)) [1..3]

map (\x -> x+1) $ [1..3]
filter isEven . map (\x -> x+1) [1..3]

-- Haskell v

quicksort :: Ord a => [a] -> [a]
quicksort [] = []
quicksort [x] = [x]
quicksort (pivot):xs = quicksort smaller ++ [pivot] ++ quicksort greater
    where
        smaller = filter (\x -> x <= pivot) xs
        greater = filter (\x -> pivot < x) xs

x: Heltal
x: x even and x * 3 = 12
returns 4 : bot

cons [] != array []
dict {}

Hashable a => map[a]b

array = array.


do =
    let input <- IO.GetInput() >>= toLowerCase
    imperative..


newtype Array

Bounded(min, max)

@
type Dimension = Integer | (Integer, Integer) | ...

Array.length ! min:0, max:number_of_countries

Compile time bounded:
bounded_integer @ Integer :: min:0, max:200, not:[42]

newtype Slot = {
    length Integer
    room Room
}
newtype Task = {
    length: Integer
    name: String
}

function schedule(slots: [Slot], tasks: [Task]) returns ... {
    -- Disjoint

    tasks uses multiple slots, never use same slot more than once, use same room for same task
    | tasks:(t:ts) = assign(t, slots)

    assign(task)
    for each task in tasks ->
        there exists multiple continuous slots with length longer than task length, in the same room

    return map{Task} this.is.a.list -> [Slot]
}

function Array.new(dimension:Dimension) returns Array {
    -- Lol bytecode.
}

Array.new(dimension:(3, 2))




function at(collection:Array, index:Integer) returns a {
    @ index >= 0
    ! index < collection.length-1
    ! index < Array.length(arr:collection)-1
    return Array. collection
}


