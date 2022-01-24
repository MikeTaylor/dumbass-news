all: nouns

nouns: sorted-nouns sorted-verbs extra-verbs non-nouns extra-nouns
	comm -23 sorted-nouns sorted-verbs | \
		comm -23 - extra-verbs | \
		comm -23 - non-nouns | \
		cat - extra-nouns | sort > $@

sorted-nouns: original/5_1_all_rank_noun.txt Makefile
	sed 's/^	\(.*\)	.*/\1/' $< | sort > $@

sorted-verbs: original/5_2_all_rank_verb.txt Makefile
	sed 's/^	\(.*\)	.*/\1/' $< | sort > $@

clean:
	rm -f sorted-nouns sorted-verbs nouns

