# go-scribe

Tool for searching long-form audio for desired keywords/concepts. Motivation for the project was to be able to search recorded
lectures for keywords and concepts instead of having to manually search through the entire video.

 Written in golang and uses Google Speech-To-Text asynchronous API to be able to handle videos longer than an hour within a few minutes. 
 
 ### Sample Usage
 
 ``` 
 go run cmd/googlestt/main.go gs://[bucket]/example_audio/[audio.wav] > example_transcripts/Lecture.txt
 go run cmd/googlestt_search/main.go example_transcripts/Lecture.txt
 Enter comma delimited list of keywords: linear,model
 Keyword:  model     Times (s):  [111. 130. 133. 137. 300. 301. 329. 743. 951. 1035. 1644. 2396. 2749. 2779. 2802. 2836. 2920. 3010. 3040. 3042. 3078. 3097. 3111. 3114. 3115. 3118. 3385. 3448. 3496. 3501. 4103. 4113. 4115. 4118. 4168. 4435. 4503.]
 Keyword:  linear     Times (s):  [1350. 1355. 1357. 1470. 1472. 1499. 1501. 1509. 1563. 1580. 1710. 1737. 1794. 1837. 1923. 2080. 2795. 3371. 3375. 3375. 3377. 3406. 4147. 4152. 4443.]
