const userAnswers = new Map; // 사용자의 답변 저장 객체
const correctAnswers = {q1:2};
var userTextAnswer;

// 선택지 선택 함수
function selectOption(questionId, optionValue) {
  userAnswers[questionId] = optionValue;
  const options = document.querySelectorAll(`#${questionId} .option`);
  
  options.forEach((option, index) => {
    option.classList.remove('selected'); // 모든 선택지에서 선택 표시 제거
    if (index + 1 === optionValue) {
      option.classList.add('selected'); // 선택된 옵션만 표시
    }
  });
}

// 시험 제출 함수
function submitExam() {
  textAnswer();

  let resultMessage = '선택된 답변:\n';
  for (const question in userAnswers) {
    resultMessage += `${question}: ${userAnswers[question]}번 \n`;
  }

  resultMessage += `마지막 문항 : ` + userTextAnswer;

  // 결과 영역에 표시
  document.getElementById('result').innerText = resultMessage;
  console.log('사용자 답변:', userAnswers);

  result()
}

// 주관식
function textAnswer() {
  var answer = document.getElementById("q_ex").value;
  console.log(answer);

  userTextAnswer = answer;
}

function result() {
  let score = 0;
  for (const question in correctAnswers) {
    if(userAnswers[question] === correctAnswers[question]) {
      score += 10;
    }
  }

   // 결과 표시
   document.getElementById("quiz-section").classList.add("hidden");
   document.getElementById("result-section").classList.remove("hidden");

   const resultText = `총 점수: ${score}점`;
   document.getElementById("result").innerText = resultText;
}