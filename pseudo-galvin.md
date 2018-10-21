
class Number a {
    also Comparable, Enumerable
    (+) :: (expr1:a, expr2:a) returns result:a
    (-) :: (expr1:a, expr2:a) returns result:a
    (*) :: (expr1:a, expr2:a) returns result:a

    negate :: (expr:a) returns neg:a
    abs :: (expr:a) returns value:a
    ...
}
() = type struct = func()
xs = [1,2,3,4,5,6,7,8,9]

filter(predicate:(even), collections:(xs))

a: some value
a -1
if a >= 0
list[a]

-- Derives is haskell magic for auto instantiate.
newtype Integer (derives Show)
-- Pseudo 
-- Do class instances need to have the same named parameters.
instance Integer as Number {
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
type Bool = True | False
type Bool = :true | :false
type Bool = true | false

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

