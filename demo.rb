class Level < ActiveRecord::Base
  has_many :lessons

  def to_s
    "#{title} -- #{description} -- #{points}"
  end
end

class Lesson < ActiveRecord::Base
  has_many :questions
  belongs_to :level
end

class Question < ActiveRecord::Base
  belongs_to :lesson
end