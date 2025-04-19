// import { useState } from 'react'
// import { useForm, SubmitHandler } from 'react-hook-form'
// import { ButtonPush } from '@/shared/ui/ButtonPush'

// type FormValues = {
//   course: string
//   semester: string
//   specialty: string
//   graduation: string
//   groupNumber: string
// }

// export function ButtonPushDisciplines() {
//   const [isOpen, setIsOpen] = useState(false)
//   const [alreadyExistsError, setAlreadyExistsError] = useState(false)

//   const {
//     register,
//     handleSubmit,
//     formState: { errors, isDirty, isValid },
//     reset,
//     watch,
//     trigger,
//   } = useForm<FormValues>({
//     defaultValues: {
//       course: '',
//       semester: '',
//       specialty: '',
//       graduation: '',
//       groupNumber: '',
//     },
//     mode: 'onChange',
//     reValidateMode: 'onChange',
//   })

//   const onSubmit: SubmitHandler<FormValues> = (data) => {
//     // setAlreadyExistsError(true)

//     setAlreadyExistsError(false)

//     console.log('Saved data:', data)
//     reset()
//     setIsOpen(false)
//   }

//   const handleCancel = () => {
//     reset()
//     setIsOpen(false)
//   }

//   return (
//     <>
//       <ButtonPush onClick={() => setIsOpen((prev) => !prev)}>Добавить предмет</ButtonPush>
//       <ModalWrapper isOpen={isOpen} onClose={handleCancel}>
//         <div className="max-w-md mx-auto p-6 bg-white rounded shadow-sm">
//           <h2 className="text-xl font-semibold mb-4">Добавление группы и предметов</h2>

//           <form onSubmit={handleSubmit(onSubmit)} className="space-y-4">
//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Курс (1/4)</label>
//               <input
//                 type="number"
//                 placeholder="1"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.course ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.course && watch('course') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('course', {
//                   required: 'Курс обязателен',
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (Number.isNaN(num)) {
//                       return 'Курс должен быть числом'
//                     }
//                     if (num < 1 || num > 4) {
//                       return 'Курс должен быть между 1 и 4'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('course')}
//               />
//               {errors.course && <p className="text-red-500 text-sm mt-1">{errors.course.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Семестр (1/2)</label>
//               <input
//                 type="number"
//                 placeholder="2"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.semester ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.semester && watch('semester') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('semester', {
//                   required: 'Семестр обязателен',
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (Number.isNaN(num)) {
//                       return 'Семестр должен быть числом'
//                     }
//                     if (num < 1 || num > 2) {
//                       return 'Семестр должен быть 1 или 2'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('semester')}
//               />
//               {errors.semester && <p className="text-red-500 text-sm mt-1">{errors.semester.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Специальность</label>
//               <input
//                 type="text"
//                 placeholder="ИСП"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.specialty ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.specialty && watch('specialty') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('specialty', {
//                   required: 'Специальность обязательна',
//                   validate: (val) => {
//                     const trimmed = val.trim()
//                     if (!trimmed) {
//                       return 'Специальность не может быть пустой'
//                     }
//                     if (trimmed.length < 2) {
//                       return 'Минимум 2 символа'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('specialty')}
//               />
//               {errors.specialty && <p className="text-red-500 text-sm mt-1">{errors.specialty.message}</p>}
//             </div>

//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Выпускники (9/11)</label>
//               <input
//                 type="text"
//                 placeholder="9"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.graduation ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.graduation && watch('graduation') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('graduation', {
//                   required: 'Укажите выпуск (9 или 11)',
//                   validate: (val) =>
//                     val === '9' || val === '11' ? true : 'Выпускной класс должен быть либо 9, либо 11',
//                 })}
//                 onBlur={() => trigger('graduation')}
//               />
//               {errors.graduation && <p className="text-red-500 text-sm mt-1">{errors.graduation.message}</p>}
//             </div>

//             {/* Номер группы */}
//             <div>
//               <label className="block text-sm font-medium text-gray-700 mb-1">Номер группы</label>
//               <input
//                 type="number"
//                 placeholder="45"
//                 className={`
//                   border rounded px-3 py-2 w-full focus:outline-none
//                   focus:ring-1 focus:ring-blue-500
//                   ${errors.groupNumber ? 'border-red-500' : 'border-gray-300'}
//                   ${!errors.groupNumber && watch('groupNumber') ? 'focus:border-green-500' : ''}
//                 `}
//                 {...register('groupNumber', {
//                   required: 'Номер группы обязателен',
//                   pattern: {
//                     value: /^\d+$/,
//                     message: 'Номер группы должен содержать только цифры',
//                   },
//                   validate: (val) => {
//                     const num = Number(val)
//                     if (num > 999) {
//                       return 'Номер группы не может быть больше 999 (пример)'
//                     }
//                     return true
//                   },
//                 })}
//                 onBlur={() => trigger('groupNumber')}
//               />
//               {errors.groupNumber && <p className="text-red-500 text-sm mt-1">{errors.groupNumber.message}</p>}
//             </div>

//             {alreadyExistsError && <p className="text-red-500 text-sm">Группа уже существует</p>}

//             <div className="flex items-center gap-4 pt-2">
//               <button
//                 type="submit"
//                 disabled={!isValid}
//                 className={`
//                   font-semibold py-2 px-4 rounded text-white
//                   ${isValid ? 'bg-green-500 hover:bg-green-600' : 'bg-gray-400 cursor-not-allowed'}
//                 `}
//               >
//                 Сохранить
//               </button>
//               <button
//                 type="button"
//                 onClick={handleCancel}
//                 className="
//                   bg-gray-300 hover:bg-gray-400
//                   text-black font-semibold
//                   py-2 px-4 rounded
//                 "
//               >
//                 Отмена
//               </button>
//             </div>
//           </form>
//         </div>
//       </ModalWrapper>
//     </>
//   )
// }
