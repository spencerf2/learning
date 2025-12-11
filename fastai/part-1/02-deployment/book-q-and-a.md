# Chapter 2: From Model to Production - Questionnaire

Questions from the fastai book, Chapter 2.

## Questions

1. Provide an example of where the bear classification model might work poorly in
   production, due to structural or style differences in the training data.

2. Where do text models currently have a major deficiency?

3. What are possible negative societal implications of text generation models?

4. In situations where a model might make mistakes, and those mistakes could be harmful,
   what is a good alternative to automating a process?

   Answers Before Starting Lesson:
      - Think this will come down to having a human in the loop.

   Answers After Completing Lesson:
      - [Answer here]

5. What kind of tabular data is deep learning particularly good at?

6. What's a key downside of directly using a deep learning model for recommendation
   systems?

   Answers Before Starting Lesson:
      - Recommendations are similar to what the user is already aware of leading to
        an echo chamber.

   Answers After Completing Lesson:
      - [Answer here]

7. What are the steps of the Drivetrain Approach?

8. How do the steps of the Drivetrain Approach map to a recommendation system?

9. Create an image recognition model using data you curate, and deploy it on the
   web.

10. What is `DataLoaders`?

11. What four things do we need to tell fastai to create `DataLoaders`?

12. What does the `splitter` parameter to `DataBlock` do?

13. How do we ensure a random split always gives the same validation set?

14. What letters are often used to signify the independent and dependent variables?

   Answers Before Starting Lesson:
      - `x` is independent (our given) `y` is dependent (what we're aiming to predict).

   Answers After Completing Lesson:
      - [Answer here]

15. What's the difference between the crop, pad, and squish resize approaches? When
    might you choose one over the others?

   Answers Before Starting Lesson:
      - Crop: cut out whatever doesn't fit within the size shape chosen.
      - Pad: add solid color to fill blank space and make the image match the
        selected dimensions.
      - Squish: force the image into the chosen dimensions.

   Answers After Completing Lesson:
      - [Answer here]

16. What is data augmentation? Why is it needed?

   Answers Before Starting Lesson:
      - Adjusting your data so that the model can use it as effectively as possible.

   Answers After Completing Lesson:
      - [Answer here]

17. What is the difference between `item_tfms` and `batch_tfms`?

18. What is a confusion matrix?

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
      - [Answer here]

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
