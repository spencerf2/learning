# Chapter 2: From Model to Production - Questionnaire

<https://github.com/fastai/fastbook/blob/master/02_production.ipynb>

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

   - There are not many pre-trained that can easily be updated for a specific
     use case.
   - They write things that are believable but factually false and it's hard
     to tell.

3. What are possible negative societal implications of text generation models?

   - They're biased based on training data. For example, looking for "healthy skin"
     is likely to return web results that'd more accurately match "young white
     woman touching her face."
   - Also, they write well and are believable, but factually false which could lead
     to an increase in accidental or intentional "fake news" style disinformation.

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

   - Define your objective
   - Figure out what actions you can take to work toward that objective
   - Define what data you have or can collect
   - Use various pieces of data to design a collection of models that can get you
     to your objective while at the same time considering what data you'd like
     to continue collecting going forward.

8. How do the steps of the Drivetrain Approach map to a recommendation system?

   - You define your objective, recommending products to people that they
     would buy if they were recommended, but wouldn't buy if they were not
     recommended.
   - You consider what data you have--other users' purchase history and your
     catalogue of products.
   - You use models to create accurate semantic or categorical
     representations of data points.
   - Then feed all of that tabular data into two different deep learning models,
     one for predicting the probability that the person will buy if shown the
     recommendation. The other predicting the probability that the person would
     buy without being shown the recommendation. You take the difference between
     the two outputs as the loss to then use for future training.
   - You optimize for a high difference between the two predictions, which
     means the recommendation had a positive effect on the person purchasing.

9. Create an image recognition model using data you curate, and deploy it on the
   web.

10. What is `DataLoaders`?

    - A class used in preparing data the model will use for training.
    - It's composed of `DataLoader`s which can be used for training
      or validation.

11. What four things do we need to tell fastai to create `DataLoaders`?

    - What kind of data we have.
      - including what the independent and dependent data points are in the data.
    - Where to find the items.
    - How to label the data.
    - How to split the data between training and validation sets.

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
      - Adjusting the brightness, contrast, colors, aspect ratio, etc. of the training data
        so that each image is viewed in a variety of ways each epoch, one way per image
        per epoch. This results in the model seeing a greater variety of images and
        learning that each should be classified the same way (whatever the label is
        on that particular image). This enhances the model's accuracy on the validation set.

17. What is the difference between `item_tfms` and `batch_tfms`?

    - The first is how to transform each item. It's handed over to CPU one item at a time.
    - The second is about transforming a group of items. It's used by a GPU thereby
      enabling concurrency in the batch.

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
      - Both the architecture (model) and its parameters. This keeps the model you trained
        together with its parameters in one, `.pkl` file ensuring that you don't have to
        manually link architecture with parameters and risk mixing up which params go
        with which model. It also makes deployment easier.

20. What is it called when we use a model for getting predictions, instead of training?

    - Inference.

21. What are IPython widgets?

    - Tools that enable you to add functionality to Jupyter Notebooks. For example, add
      a GUI for relabeling / deleting bad data the model is being trained on. Also,
      IPython widgets can help you deploy your Jupyter Notebook so it's a live website.

22. When might you want to use CPU for deployment? When might GPU be better?

   Answers Before Starting Lesson:
      - CPU could be fine if that's all your able to access due to demand crowding out
        access to a GPU. Also fine if you're working with a smaller or light weight model.
      - GPU better for larger models and models with many params. Better when you need
        concurrency. I believe that concurrency would be related to evaluating parts of the
        model against different sections of the params, but am not really sure.

   Answers After Completing Lesson:
      - GPU isn't really needed if you're only training a few epochs.
      - If you need to concurrently evaluate requests or process requests in batches due
        to high volume, then GPU is helpful, though costs more.
      - If you just evaluate requests sequentially, and there isn't a lot of
        computation required in the evaluations, then CPU is fine and less expensive.

23. What are the downsides of deploying your app to a server, instead of to a client
    (or edge) device such as a phone or PC?

   Answers Before Starting Lesson:
      - Server means one central location so higher latency.
      - At the edge/client would mean it's closer to users so faster responses.

   Answers After Completing Lesson:
      - Latency
      - Privacy--users worrying about their data being collected and trained on.

24. What are three examples of problems that could occur when rolling out a bear
    warning system in practice?

    - The model doesn't do a good job of recognizing bears in videos because it was
      trained on images.
    - The model needs human intervention with the video data to adjust to videos
      instead of images.
    - The model doesn't identify bears until they're close to the camera or doesn't
      detect when it's night time.

25. What is "out-of-domain data"?

    - When the data that the model makes predictions on is different than the
      data the model was trained on. For example, a model trained to detect bears
      in images taken from the internet being used to detect bears in videos
      which are recorded at all hours of the night and which feature bears hidden
      much different angles than the clear pictures the model was trained on.

26. What is "domain shift"?

    - When the data that the model sees and makes predictions upon differs from
      the data that the model was trained to make predictions upon. For example,
      an insurance company predicting risk may, over time, evaluate the risk of
      customers the same way it did years ago, when really the types of customers
      it now servers have substantially different risk profiles than when their
      model used as part of the risk evaluation was trained.

27. What are the three steps in the deployment process?

   Answers Before Starting Lesson:
      - Train the model, export the model, host/deploy the model so it's accessible.

   Answers After Completing Lesson:
      - Have a human review all the model's inferences. The inferences are not
        used at all to do anything.
      - Have a limited scope deployment. Limit time and location. Maybe just 1
        location. Still have a human involved, but now the model's inferences are
        used to drive actions, but are also human reviewed for accuracy.
      - Gradually roll the model out to more locations and for a longer duration,
        slowly involving human supervision less and less. Good reporting features
        are key to ensure problems get detected.
