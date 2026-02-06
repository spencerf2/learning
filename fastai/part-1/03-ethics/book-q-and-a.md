# Chapter 3: Data Ethics - Questionnaire

<https://github.com/fastai/fastbook/blob/master/03_ethics.ipynb>

Questions from the fastai book, Chapter 3.

## Questions

1. Does ethics provide a list of "right answers"?

   Answer: No. It provides a few frameworks for thinking through a problem
   or situation and evaluating it critically to get at the ethical
   considerations.

2. How can working with people of different backgrounds help when considering
   ethical questions?

   Answer: They think differently due to their different backgrounds
   reducing the surface area for running into blind spots.

   For example, if everyone on the team is of the same race, gender,
   ethnicity, etc., then they're more likely to have the same expectations
   when they go to search Google or try out the projects their working on.
   Someone of a different background may search differently or try out
   the project in a different way leading to discoveries that otherwise
   may have been ignored. Given the power of deep learning and training
   models, small things unnoticed can have massive impacts on society.

3. What was the role of IBM in Nazi Germany? Why did the company participate
   as it did? Why did the workers participate?

   Answer: IBM's goal was to maximize profits and best service its
   customers. They made machines which could be used to provide accurate
   record keeping. That led them to doing the best they could do at
   identifying people of various races, ethnicities, etc., leading
   to a much deeper categorization of people on the German census.

   The workers felt they were doing their job, and focused on that,
   rather than what the work they were doing was used for. One way this
   can happen is that a company many not tell everyone working on a
   given project how their work is connected with other members of the
   team. Everyone working in silos thinks their work is doing good things.
   They don't know what others are working on and how combining it all
   together may be doing something truly catastrophic for society.

4. What was the role of the first person jailed in the Volkswagen diesel
   scandal?

   Answer: (Had to look this up in the book) One of the engineers.
   It wasn't the manager of the project or an executive as one might
   expect. It was actually an engineer working on the project who just
   did what he was told.

   Side note: I'm glad to be reminded of this through this question.

5. What was the problem with a database of suspected gang members maintained
   by California law enforcement officials?

   Answer: (Had to look this up in the book) It was full of errors, one of
   which was that it had 42 babies, which were added to the database when
   they were less than 1 year old--and those babies had "admitted to being
   gang members". There was no process in place to correct errors or remove
   people once they were added.

6. Why did YouTube's recommendation algorithm recommend videos of partially
   clothed children to pedophiles, even though no employee at Google had
   programmed this feature?

   Answer: It noticed that users who were interested in one home video
   containing partially clothed children was also interesting to other
   users interested in videos of partially clothed children. That led to a
   feedback loop where if someone started watching one home video, others
   were recommended. When that thread continues to have more people
   pull on it, it leads YouTube to get really good at recommending such
   a feed to such users. It wasn't intentional for that to happen, but
   it did.

   One thing that can be done to mitigate this is to not include videos
   containing children in recommendations. If YouTube just did not
   recommend videos of children or videos of partially clothed children
   to users, then a feedback loop like that couldn't have been created.

7. What are the problems with the centrality of metrics?

   Answer: (Had to look this up in the book) Metrics are just a number
   devoid of humanness. Architectures will optimize the metric at all
   costs leading to unexpected outcomes--as in the case where YouTube
   curated pedophile playlists.

8. Why did Meetup.com not include gender in its recommendation system for
   tech meetups?

   Answer: They realized that it would further bias the recommendations
   so that fewer women were recommended tech meetups since there were
   significantly more men showing an interest in them.

   This choice helped protect women's interests and ensure that they
   did not miss out on the opportunity to discover tech meetups,
   which would have further exacerbated the existing bias or pattern
   inherent in the "gendered" interests.

9. What are the six types of bias in machine learning, according to Suresh
   and Guttag?

   Answer: (Had to look this up in the book) They are:
   - Historical Bias
     - The data the model is built on or its assumptions already include
       bias through the sheer fact of their humanness and not looking out
       for bias during selection (though even if they had, I understand
       from this chapter that there'd still be some bias present, though
       we should always look at how we can minimize it to avoid unexpected
       negative outcomes.)
   - Representation Bias
     - When the data itself represents things in an unequal way. For example
       data that's based on one geographic region, when it'll be used in
       other regions. It'll skew toward the region it was trained on even
       if that doesn't match the new geography well.
   - Measurement Bias
     - Measuring the wrong thing, or in the wrong way, or inappropriately
       incorporating a measurement into the model.
   - Evaluation Bias
     - When the benchmark data used for evaluation/testing is different than the
       population the data will actually operate upon. Thinking of LLMs, it's
       the type of bias present when a company tests their model against LLM
       tests, such as humanity's last exam, and the model does well, but in
       actuality, the way users will use the model isn't really like that
       test and so the model underperforms for the real use-case.
   - Aggregation Bias
     - When a model doesn't include the appropriate interactions among the
       data collected. Or when a model doesn't include all factors needed
       when it's put together. Or when relationships are compiled in a way
       that's too simplistic and doesn't actually model reality.
     - For example, diabetes depends on gender and ethnicity, but the current
       way of diagnosing diabetes is based on identifying HbA1c levels which
       differ between genders and ethnicities leading to inaccurate diagnoses.
   - Deployment Bias
     - When the model was intended to solve one problem, but the way it's
       deployed utilizes it for a different problem than intended.

10. Give two examples of historical race bias in the US.

    Answer: (Had to look this up in the book)
    1. An all white jury was 16% more likely to convict a defendant when they
       were black, but adding 1 black member to the jury neutralized this.
    2. Responding to Craigslist ads with a black name resulted in fewer
       responses than a white name.

11. Where are most images in ImageNet from?

    Answer: (Had to look this up in the book) The US and other Western
    countries. Research has found that models trained on that data perform
    worse at recognizing basic household items from lower income countries.

12. In the paper "Does Machine Learning Automate Moral Hazard and Error" why
    is sinusitis found to be predictive of a stroke?

    Answer: (Had to look this up in the book) Because of measurement bias
    in the data. People who go to the doctor for accidental injuries and
    other such issues are likely to go to the doctor when something major
    like a stroke happens. That means people with sinusitis are more likely
    to be associated with a stroke simply because they're the type of people
    who go to the doctor period.

13. What is representation bias?

    Answer: (Had to look this up in the book) When there's already bias in
    the data (fewer women are surgeons than men) and a model--usually a
    simple one--amplifies that bias (for ex. predicting even fewer women
    to be surgeons than the original data showed.)

14. How are machines and people different, in terms of their use for making
    decisions?

    Answer: Sometimes the most optimized decision a machine can make is
    one that humans would immediately recognize as not acceptable, such
    as turning elderly people into fertilizer (which a paper discussed
    in this chapter identified as an optimal machine choice that no
    human would accept).

    Also, humans tend to think of models as objective and error free
    when in reality they can actually amplify bias and lead to awful
    decisions and outcomes for society. The book points out that the
    privileged are often processed by people while the poor get
    processed by models since it's cheaper.

15. Is disinformation the same as "fake news"?

    Answer: No. Fake news may be unintentional. For example reporting
    too early for the facts to be clear, or reporting the current
    prevailing idea that turns out to be false. Disinformation is
    intentional and often contains parts of the truth that are
    intentionally twisted to be just true enough to be convincing
    but are ultimately false and pushing an agenda.

16. Why is disinformation through auto-generated text a particularly
    significant issue?

    Answer: Auto-generated text is quick to create and inexpensive making
    it easy to produce at scale. Additionally, it's based on things
    that were likely true, or were realistic, and therefore when the
    model generates the text, it's really tricky to spot where it goes
    from being truthy to false. It's also well written and confident
    sounding.

17. What are the five ethical lenses described by the Markkula Center?

    Answer: (Had to look this up in the book)
    1. The rights approach.
       1. Which option best represents the rights of all those who have
          a stake?
    2. The justice approach.
       1. Which option treats people equally or proportionally?
    3. The utilitarian approach.
       1. Which option will produce the most good and the least harm?
    4. The common good approach.
       1. Which option best serves the community as a whole, not just
          some members?
    5. The virtue approach.
       1. Which option leads me to act as the sort of person I want to
          be?

18. Where is policy an appropriate tool for addressing data ethics issues?

    Answer: Where it levels the playing field so that everyone must factor
    it in and it's not something one group does because they're willing
    to sacrifice profits while another group is profit seeking and ok
    with passing the negative externalities onto society.
