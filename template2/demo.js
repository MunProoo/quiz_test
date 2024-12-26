document.addEventListener('DOMContentLoaded', () => {
    // const options = document.querySelectorAll('.option');
    const submitBtn = document.querySelector('.submit-btn');
    const quizDiv = document.querySelector('#quiz');
    const resultsDiv = document.querySelector('.results');
    var userName;
    
    const answers = {
        1: 2,
        2: 4,
        3: 4,
        4: 1,
        5: 1,
        6: 1,
        7: 1,
        8: 1,
        9: 1,
        10: 1
    };

    // 모든 질문에 대해 처리
    document.querySelectorAll('.question').forEach((question) => {
        // 각 질문의 모든 옵션에 대해 이벤트 리스너 추가
        question.querySelectorAll('.option').forEach(option => {
            option.addEventListener('click', () => {
                // 현재 질문의 다른 옵션들의 선택 상태만 초기화
                question.querySelectorAll('.option').forEach(opt => {
                    opt.classList.remove('selected');
                });
                // 클릭한 옵션 선택
                option.classList.add('selected');
            });
        });
    });

    // options.forEach(option => {
    //     option.addEventListener('click', () => {
    //         const parentQuestion = option.closest('.question');
    //         parentQuestion.querySelectorAll('.option').forEach(opt => {
    //             opt.classList.remove('selected');
    //         });
    //         option.classList.add('selected');
    //     });
    // });

    submitBtn.addEventListener('click', () => {
        userName = document.getElementById("userName").value;
        var popularName = document.getElementById('q_ex').value;

        if(!userName) {
            alert("이름을 입력해주세요");
            document.getElementById("userName").focus();
            return;
        } else if(!popularName) {
            alert("좋아하는 사람을 꼭 입력해주세요. 제 이름이라도..!");
            document.getElementById('q_ex').focus();
            return
        }

        let score = 0;
        document.querySelectorAll('.question').forEach((question, index) => {
            const selectedOption = question.querySelector('.option.selected');
            if (selectedOption) {
                const questionNumber = index + 1;
                const selectedValue = parseInt(selectedOption.dataset.value);
                if (answers[questionNumber] === selectedValue) {
                    score += 10;
                }
            }
        });

        document.getElementById('name').textContent = userName + "님의 "
        document.getElementById('score').textContent = score;
        quizDiv.style.display = 'none';
        resultsDiv.style.display = 'block';


        // // 메일로 결과 전송
        // async function sendEmailResult() {
        //     try {
        //         const response = await fetch('/submit-result', {
        //             method: "POST",
        //             headers: {
        //                 'Content-Type': 'application/json',
        //             },
        //             body : JSON.stringify({
        //                 name : userName,
        //                 score : score,
        //                 popularName : popularName,
        //                 timestamp: new Date().toISOString()
        //             })
        //         });
    
        //         if (!response.ok) {
        //             throw new Error('이메일 전송 실패');
        //         }
            
        //     } catch (error) {
        //         console.error('Error:', error);
        //     }
        // }
        // sendEmailResult();
        
    });
});

function retry() {
    var quizDiv = document.querySelector('#quiz');
    var resultsDiv = document.querySelector('.results');
    quizDiv.style.display = 'block';
    resultsDiv.style.display = 'none';

    document.querySelectorAll('.option').forEach((option) => {
        option.classList.remove('selected');
    });

    document.getElementById('q_ex').value = "";
}