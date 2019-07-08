
# Various implementations of TFIDF

https://en.wikipedia.org/wiki/Tf%E2%80%93idf

function is(Comparable a => x: a, y: a) returns Boolean {
	return x = y
}

-- binary weighting scheme
function tf(term: String, document: String) returns Boolean {
	return document.contains(term)
}


val (line_number, line) = (1, "asdf")

0x

-- raw count weighting scheme
function occurences(term: String, document: String) returns Integer {
	return count . filter(predicate: is(x: term)) . words(string: document)
}
tf = occurences

-- raw count weighting scheme
function tf(term: String, document: String) returns Real {
	let words = strings.words(string: document)
	let occurences = count . filter(predicate: is(x: term)) . words


	let occurences = count . filter(predicate: is(x: term)) . words
	= 
	let occurences = count . filter & is(x: term) . words

	length . words


	return 
}


-- The inverse document frequency is a measure of how much information the word
-- provides, i.e., if it's common or rare across all documents. It is the
-- logarithmically scaled inverse fraction of the documents that contain the word
-- (obtained by dividing the total number of documents by the number of documents
-- containing the term, and then taking the logarithm of that quotient)

function idf(term: String, documents: [String]) returns Real {
	let num_documents: length documents

	let in_document = \document -> tf term document != 0
	let term_occurences = length . filter in_document documents

	-- We add 1 to the denominator to prevent unwanted division by zero.
	let quotient = num_documents / (1 + term_occurences)

	return math.log(quotient)
}
