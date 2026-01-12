# Chapter 2: From Model to Production - Questionnaire

Questions from the fastai book, Chapter 2.

## Questions

1. Provide an example of where the bear classification model might work poorly in
   production, due to structural or style differences in the training data.

   - Because it was trained on pictures found via web search, when moving to prod
     it'll be seeing video rather than pictures, and the video it sees will be
     very different. Web pictures that humans took often photo bears from a
     photogenic angle ("this is a bear"). In the real world, the video will see
     bears hidden behind bushes, at odd angles (back of bear, just a bear paw on
     screen, etc.), and with odd lighting (bear in the night).
   - All this means when we roll to prod, we should go slow, with the park ranger
     manually reviewing model predictions. We can retrain with their corrections
     then try rolling out the model to a single park for a week (location/time
     bounded trial). It can be trained slowly over time on the real video data
     before being truly rolled out to full use in all national parks.

2. Where do text models currently have a major deficiency?

3. What are possible negative societal implications of text generation models?

   - They're biased based on training data. For example, looking for "healthy skin"
     is likely to return web results that'd more accurately match "young white
     woman touching her face."

4. In situations where a model might make mistakes, and those mistakes could be harmful,
   what is a good alternative to automating a process?

   Answers Before Starting Lesson:
      - Think this will come down to having a human in the loop.

   Answers After Completing Lesson:
      - Slowly rolling out, starting with manual review by a human,
        then removing the human more and more as the training set
        used for the model more accurately captures the reality
        of where the model will be deployed. (See Question 1's bear example.)

5. What kind of tabular data is deep learning particularly good at?

   - Categorical and where a column represents something semantic.
   - The model can handle vast arrays of data.
   - Good example is Amazon, which has a massive matrix with every product on one
     side and every customer on the other axis. Deep learning is great at using
     such data to predict what another customer will be interested in purchasing
     based off what similar customers purchased.

6. What's a key downside of directly using a deep learning model for recommendation
   systems?

   Answers Before Starting Lesson:
      - Recommendations are similar to what the user is already aware of leading to
        an echo chamber.

   Answers After Completing Lesson:
      - The model can only make recommendations based on what the user may already
        be aware of rather than suggesting novel items they may not be aware of.
      - Amazon was used as an example. The model could recommend products the customer
        already owns. For example, if the customer purchased the first book in a series
        the model may recommend another book in that series. That might be relevant to
        the customer and something they would purchase, but that might not result in
        them buying or in a helpful recommendation.

7. What are the steps of the Drivetrain Approach?

8. How do the steps of the Drivetrain Approach map to a recommendation system?

9. Create an image recognition model using data you curate, and deploy it on the
   web.

10. What is `DataLoaders`?

    - A class used in preparing data the model will use for training.

11. What four things do we need to tell fastai to create `DataLoaders`?

    - What kind of data we have.
    - Where to find the items.

12. What does the `splitter` parameter to `DataBlock` do?

    - Specify how to split the data into training and validation sets.

13. How do we ensure a random split always gives the same validation set?

    - By specifying the random seed. When the seed is the same, the "random
      outcome" is the same.

14. What letters are often used to signify the independent and dependent variables?

   Answers Before Starting Lesson:
      - `x` is independent (our given) `y` is dependent (what we're aiming to predict).

   Answers After Completing Lesson:
      - Before answer was correct.

15. What's the difference between the crop, pad, and squish resize approaches? When
    might you choose one over the others?

   Answers Before Starting Lesson:
      - Crop: cut out whatever doesn't fit within the size shape chosen.
      - Pad: add solid color to fill blank space and make the image match the
        selected dimensions.
      - Squish: force the image into the chosen dimensions.

   Answers After Completing Lesson:
      - Crop & Pad keep the correct aspect ratio. Crop loses part of the picture.
      - Squish keeps everything in the pic but you lose the aspect ratio.
      - RandomResizedCrop with aug_transforms is most like what you'll want to do
        in practice. It ensures you get different parts of an image, and when you're
        training over many epochs (5+) that helps increase accuracy.

16. What is data augmentation? Why is it needed?

   Answers Before Starting Lesson:
      - Adjusting your data so that the model can use it as effectively as possible.

   Answers After Completing Lesson:
      - [Answer here]

17. What is the difference between `item_tfms` and `batch_tfms`?

18. What is a confusion matrix?

    Answers After Completing Lesson:
       - A graph with predictions on one axis and actual on the other.
         Shows us which classifications we get wrong.
       - Can be used with `plot_top_losses` to see where the accuracy is
         worst (wrong confident prediction / correct uncertain prediction).
       - `ImageClassifierCleaner()` can be used to clean (reclassify / delete)
         data.

19. What does `export` save?

   Answers Before Starting Lesson:
      - I'm guessing that it means exporting a pickle file of the model you've trained.

   Answers After Completing Lesson:
      - [Answer here]

20. What is it called when we use a model for getting predictions, instead of training?

21. What are IPython widgets?

22. When might you want to use CPU for deployment? When might GPU be better?

   Answers Before Starting Lesson:
      - CPU could be fine if that's all your able to access due to demand crowding out
        access to a GPU. Also fine if you're working with a smaller or light weight model.
      - GPU better for larger models and models with many params. Better when you need
        concurrency. I believe that concurrency would be related to evaluating parts of the
        model against different sections of the params, but am not really sure.

   Answers After Completing Lesson:
      - GPU isn't really needed if you're only training a few epochs.

23. What are the downsides of deploying your app to a server, instead of to a client
    (or edge) device such as a phone or PC?

   Answers Before Starting Lesson:
      - Server means one central location so higher latency.
      - At the edge/client would mean it's closer to users so faster responses.

   Answers After Completing Lesson:
      - [Answer here]

24. What are three examples of problems that could occur when rolling out a bear
    warning system in practice?

25. What is "out-of-domain data"?

26. What is "domain shift"?

27. What are the three steps in the deployment process?

   Answers Before Starting Lesson:
      - Train the model, export the model, host/deploy the model so it's accessible.

   Answers After Completing Lesson:
      - [Answer here]
